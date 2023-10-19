package pg

import (
	"context"
	"database/sql"
	// "fmt"
	"log"
	r "regate-core/domain/repository"
)

type empresaRepo struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) r.EmpresaRepository{
	return &empresaRepo{
		Conn: conn,
	}
}
func (p *empresaRepo)GetUuidEmpresa(ctx context.Context,id int)(uuid string,err error){
	query := `select uuid from empresas where empresa_id = $1`
	err = p.Conn.QueryRowContext(ctx,query,id).Scan(&uuid)
	return 
}

func (p *empresaRepo) CreateEmpresa(ctx context.Context,d *r.Empresa)(err error){
	query := `insert into empresas(name,phone_number,email) values($1,$2,$3) returning uuid,empresa_id,created_at`
	err = p.Conn.QueryRowContext(ctx,query,d.Name,d.PhoneNumber,d.Email).Scan(&d.Uuid,&d.Id,&d.CreatedAt)
	return
}

func (p *empresaRepo) GetEmpresasByEstado(ctx context.Context,estado r.EmpresaEstado)(res []r.Empresa,err error){
	if estado == 0 {
		query := `select empresa_id,uuid,name,created_at from empresas;`
		res,err = p.fetchEmpresas(ctx,query)
		return
	}else {
		query := `select empresa_id,uuid,name,created_at from empresas where estado = $1`
		res,err = p.fetchEmpresas(ctx,query,estado)
		return
	}
}

func (p *empresaRepo) GetEmpresa(ctx context.Context,uuid string)(res r.Empresa,err error){
	query := `select empresa_id,uuid,name,estado,created_at,email,phone_number,
	address,ST_X(geog::geometry),ST_Y(geog::geometry) from empresas
	where uuid = $1`
	err = p.Conn.QueryRowContext(ctx,query,uuid).Scan(&res.Id,&res.Uuid,&res.Name,&res.Estado,&res.CreatedAt,&res.Email,
		&res.PhoneNumber,&res.Address,&res.Longitud,&res.Latidud)
	return	
}
// func (p *empresaRepo) GetDepositos(ctx context.Context,d r.DepositoFilterData,page int16,size int8) (res []r.DepositoBancario, 
// 	count int,err error) {
// 	var (
// 		empresaFilter string
// 	)	
// 	if d.EmpresaId == 0 {
// 		empresaFilter = ""
// 	}else{
// 		empresaFilter = fmt.Sprintf("where empresa_id = %s",strconv.Itoa(d.EmpresaId))
// 	}
// 	query := fmt.Sprintf(`select id,uuid,income,created_at,gloss from deposito_bancario 
// 	%s limit $1 offset $2`,empresaFilter)
// 	res, err = p.fetchEmpresas(ctx, query,size, page*int16(size))
// 	if int8(len(res)) >= size {
// 		query = fmt.Sprintf(`select count(*) from deposito_bancario %s`, empresaFilter)
// 		err = p.Conn.QueryRowContext(ctx, query).Scan(&count)
// 		if err != nil {
// 			return
// 		}
// 	} else {
// 		count = len(res)
// 	}
// 	return
// }

func (p *empresaRepo) fetchEmpresas(ctx context.Context, query string, args ...interface{}) (res []r.Empresa, err error) {
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
	res = make([]r.Empresa, 0)
	for rows.Next() {
		t := r.Empresa{}
		err = rows.Scan(
			&t.Id,
			&t.Uuid,
			&t.Name,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}