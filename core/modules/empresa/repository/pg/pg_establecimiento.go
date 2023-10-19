package pg

import (
	"context"
	"database/sql"

	// "fmt"
	"log"
	r "regate-core/domain/repository"
	util "regate-core/domain/util"
)

type establecimientoRepo struct {
	Conn *sql.DB
}


func NewEstablecimientoRepository(conn *sql.DB) r.EstablecimientoRepository{
	return &establecimientoRepo{
		Conn: conn,
	}
}
func(p *establecimientoRepo)UpdateEstadoEstablecimiento(ctx context.Context,
	id int,estado r.EstablecimientoEstado,visibility bool)(err error){
	query := `update establecimientos set estado= $1 ,visibility = $2 where establecimiento_id = $3`
	_,err = p.Conn.ExecContext(ctx,query,estado,visibility,id)
	return
}
func (p *establecimientoRepo)GetEstablecimientosByEstado(ctx context.Context,estado r.EstablecimientoEstado)(res []int,err error){
	query := `select establecimiento_id from establecimientos where estado = $1`
	res,err = p.fetchIds(ctx,query,r.EstablecimientoVerificado)
	return
}
func (p *establecimientoRepo)UpdateEstablecimientoTsv(ctx context.Context,id int)(err error){
	var (
		query string
		values []string
		establecimientoName string
	)
	query = `select c.name from instalaciones as i
	left join categorias as c on c.category_id = i.category_id
	where establecimiento_id = $1 group by c.name`
	categories,err := p.fetchCategoryName(ctx,query,id)
	if err != nil {
		return
	}
	values = categories
	query = `select name from establecimientos where establecimiento_id = $1`
	err = p.Conn.QueryRowContext(ctx,query,id).Scan(&establecimientoName)
	if err != nil {
		return
	}
	values = append(values, establecimientoName)
	query = `update establecimientos set tsv = to_tsvector($1) where establecimiento_id = $2`
	_,err = p.Conn.ExecContext(ctx,query,util.AppendTsv(values),id)
	return
}

func (p *establecimientoRepo)GetEstablecimientosEmpresa(ctx context.Context,uuid string)(res []r.Establecimiento,err error){
	query := `select establecimiento_id,uuid,name,visibility,estado,created_at
	 from establecimientos where empresa_id = (select empresa_id from empresas where uuid = $1)`
	res,err = p.fetchEstablecimientos(ctx,query,uuid)
	return 
}

func (p *establecimientoRepo) fetchEstablecimientos(ctx context.Context, query string, args ...interface{}) (res []r.Establecimiento, err error) {
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
	res = make([]r.Establecimiento, 0)
	for rows.Next() {
		t := r.Establecimiento{}
		err = rows.Scan(
			&t.Id,
			&t.Uuid,
			&t.Name,
			&t.Visibility,
			&t.Estado,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (p *establecimientoRepo) fetchCategoryName(ctx context.Context, query string, args ...interface{}) (res []string, err error) {
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
	res = make([]string, 0)
	for rows.Next() {
		var t string
		err = rows.Scan(
			&t,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (p *establecimientoRepo) fetchIds(ctx context.Context, query string, args ...interface{}) (res []int, err error) {
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
	res = make([]int, 0)
	for rows.Next() {
		var t int
		err = rows.Scan(
			&t,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}