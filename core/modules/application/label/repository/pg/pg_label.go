package repository

import (
	"context"
	"database/sql"
	r "regate-core/domain/repository"
	// "time"
)

type pgInfo struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) r.LabelRepository {
	return &pgInfo{
		Conn: conn,
	}
}


func (p *pgInfo) CreateAmenity(ctx context.Context,d r.Label)(err error){
	var query string
	if d.Id == 0 {
		query = `insert into amenities(amenity_name,thumbnail) values($1,$2)`
		_,err =p.Conn.ExecContext(ctx,query,d.Name,d.Thumnail)
		return
	}else{
		query = `update amenities set amenity_name = $1,thumbnail=$2 where amenity_id = $3 `
		_,err =p.Conn.ExecContext(ctx,query,d.Name,d.Thumnail,d.Id)
		return
	}
}

func (p *pgInfo) CreateCategory(ctx context.Context,d r.Label)(err error){
	var query string
	if d.Id == 0 {
		query = `insert into categorias(name,thumbnail) values($1,$2)`
		_,err =p.Conn.ExecContext(ctx,query,d.Name,d.Thumnail)
		return
	}else{
		query = `update categorias set name = $1,thumbnail=$2 where category_id = $3 `
		_,err =p.Conn.ExecContext(ctx,query,d.Name,d.Thumnail,d.Id)
		return
	}
}
func (p *pgInfo) CreateDeporte(ctx context.Context,d r.Label)(err error){
	var query string
	if d.Id == 0 {
		query = `insert into deportes(deporte_id) values($1)`
		_,err =p.Conn.ExecContext(ctx,query,d.Name)
		return
	}else{
		query = `update deportes set amenity_name = $1 where deporte_id = $2 `
		_,err =p.Conn.ExecContext(ctx,query,d.Name,d.Id)
		return
	}
}
func (p *pgInfo) CreateRule(ctx context.Context,d r.Label)(err error){
	var query string
	if d.Id == 0 {
		query = `insert into rules(rule) values($1)`
		_,err =p.Conn.ExecContext(ctx,query,d.Name,d.Thumnail)
		return
	}else{
		query = `update rules set rule = $1 where rule_id = $2 `
		_,err =p.Conn.ExecContext(ctx,query,d.Name,d.Id)
		return
	}
}


func (p *pgInfo)DeleteAmenity(ctx context.Context,id int)(err error){
	return 
}

func (p *pgInfo)DeleteRule(ctx context.Context,id int)(err error){
	return 
}



// func(pg *pgInfo)fetchCronJobs(ctx context.Context,query string,args ...interface{})(res []r.CronJob,err error){
// 	rows,err := pg.Conn.QueryContext(ctx,query,args...)
// 	if err != nil {
// 		return nil,err
// 	}
// 	defer func()  {
// 		rows.Close()
// 	}()
// 	res = make([]r.CronJob, 0)
// 	for rows.Next(){
// 		t := r.CronJob{}
// 		err = rows.Scan(
// 			&t.Id,
// 			&t.Name,
// 			&t.Expression,
// 			&t.Tag,
// 		)
// 		res = append(res, t)
// 	}
// 	return 
// }

