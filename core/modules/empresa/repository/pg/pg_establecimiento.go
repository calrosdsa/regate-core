package pg

import (
	"context"
	"database/sql"
	// "fmt"
	"log"
	r "regate-core/domain/repository"
)

type establecimientoRepo struct {
	Conn *sql.DB
}

func NewEstablecimientoRepository(conn *sql.DB) r.EstablecimientoRepository{
	return &establecimientoRepo{
		Conn: conn,
	}
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