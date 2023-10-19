package pg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	r "regate-core/domain/repository"
	"strconv"
)

type billingRepo struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) r.BillingRepository{
	return &billingRepo{
		Conn: conn,
	}
}
func(p *billingRepo)GetEmpresaIdByDepositoDetailId(ctx context.Context,id int)(res int,err error){
	query := `select empresa_id from deposito_bancario_detail where id = $1`
    err = p.Conn.QueryRowContext(ctx,query,id).Scan(&res)
	return
}

func(p *billingRepo)UploadComprobanteDeposito(ctx context.Context,d r.DepositoBancario)(err error){
	query := `update deposito_bancario set comprobante_url = $1,estado = $2,emition_date = current_timestamp where id = $3`
	_,err = p.Conn.ExecContext(ctx,query,d.ComprobanteUrl,r.DEPOSITO_EMITIDO,d.Id)
	return
}

func (p *billingRepo) GetDepositosEmpresa(ctx context.Context,d r.DepositoFilterData,page int16,size int8) (res []r.DepositoBancarioEmpresa, 
	count int,err error) {
	var (
		empresaFilter string
	)	
	if d.EmpresaId == 0 {
		empresaFilter = ""
	}else{
		empresaFilter = fmt.Sprintf("where empresa_id = %s",strconv.Itoa(d.EmpresaId))
	}
	query := fmt.Sprintf(`select d.id,d.uuid,d.empresa_id,e.name,d.created_at,
	(select sum(income) from deposito_bancario where parent_id = d.id)
	from deposito_bancario_detail as d
	left join empresas as e on e.empresa_id = d.empresa_id
	%s order by created_at desc limit $1 offset $2
    `,empresaFilter)
	res, err = p.fetchDepositosEmpresa(ctx, query,size, page*int16(size))
	if err != nil {
		log.Println(err)
		return
	}
	query = fmt.Sprintf(`select count(*) from deposito_bancario_detail %s`, empresaFilter)
	err = p.Conn.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return
	}

	return
}
func (p *billingRepo)CreateDeposito(ctx context.Context,empresaId int)(err error){
	var (
		query string
		parentId int
		tarifa float32
		currencyId int16
	)
	conn, err := p.Conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			conn.Rollback()
		} else {
			conn.Commit()
		}
	}()
	query = `insert into deposito_bancario_detail(empresa_id) values($1) returning id;`
	err =conn.QueryRowContext(ctx,query,empresaId).Scan(&parentId)
	if err != nil {
		log.Println(err,"-----ERROR 1")
		return 
	}
	query = `select currency_id,tarifa from empresa_settings where empresa_id = $1`
	err = conn.QueryRowContext(ctx,query,empresaId).Scan(&currencyId,&tarifa)
	if err != nil {
		log.Println(err,"ERROR 2")
		return 
	}
	query = `insert into deposito_bancario(income,gloss,establecimiento_id,tarifa,currency_id,parent_id,date_paid)
	select coalesce(sum(r.paid),0),('Pago del dia--'),e.establecimiento_id,($1),($2),($3),(current_timestamp::date - INTERVAL '1 DAY')
	from establecimientos as e 
	left join reservas as r on r.establecimiento_id = e.establecimiento_id and r.type_reserva = $4 and r.estado = $5
	and start_date::date = current_timestamp::date - INTERVAL '1 DAY'
	where e.empresa_id = $6 
	group by e.establecimiento_id returning id,establecimiento_id;`
	ids,err := p.fetchIds(conn,ctx,query,tarifa,currencyId,parentId,r.ReservaTypeApp,r.ReservaValid,empresaId)
	if err != nil {
		log.Println(err,"ERROR 2")
		return 
	}
	for _,id := range ids {
		query := `insert into facturacion_reserva(reserva_id,deposito_id) 
		select r.reserva_id,$1 from reservas as r where r.establecimiento_id = $2 and r.type_reserva = $3 and r.estado = $4
		and start_date::date = current_timestamp::date - INTERVAL '1 DAY'`
		_,err = conn.ExecContext(ctx,query,id.Id,id.EstablecimientoId,r.ReservaTypeApp,r.ReservaValid)
		if err != nil {
		    log.Println(err,"ERROR 4")
		}
	}
	return
}
func (p *billingRepo) fetchIds(conn *sql.Tx,ctx context.Context, query string, args ...interface{}) (res []r.IdsDepositoReturn, err error) {
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println(errRow)
		}
	}()
	res = make([]r.IdsDepositoReturn, 0)
	for rows.Next() {
		t := r.IdsDepositoReturn{}
		err = rows.Scan(
			&t.Id,
			&t.EstablecimientoId,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (p *billingRepo) fetchDepositosEmpresa(ctx context.Context, query string, args ...interface{}) (res []r.DepositoBancarioEmpresa, err error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println(errRow)
		}
	}()
	res = make([]r.DepositoBancarioEmpresa, 0)
	for rows.Next() {
		t := r.DepositoBancarioEmpresa{}
		err = rows.Scan(
			&t.Id,
			&t.Uuid,
			&t.EmpresaId,
			&t.EmpresaName,
			&t.CreatedAt,
			&t.TotalIncome,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (p *billingRepo) fetchDepositos(ctx context.Context, query string, args ...interface{}) (res []r.DepositoBancario, err error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println(errRow)
		}
	}()
	res = make([]r.DepositoBancario, 0)
	for rows.Next() {
		t := r.DepositoBancario{}
		err = rows.Scan(
			&t.Id,
			&t.Uuid,
			&t.Income,
			&t.CreatedAt,
			&t.Gloss,
			&t.EstablecimientoId,
			&t.DatePaid,
			&t.EstablecimientoName,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}
