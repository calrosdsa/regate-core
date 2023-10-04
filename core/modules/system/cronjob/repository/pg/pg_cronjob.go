package repository

import (
	"context"
	"database/sql"
	r "regate-core/domain/repository"
	// "time"
)

type pgCronJob struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) r.CronJobRepository {
	return &pgCronJob{
		Conn: conn,
	}
}


func (pg *pgCronJob) GetCronJobs(ctx context.Context)(res []r.CronJob,err error){
	query := `select id,name,expression,tag from cronjob`
	res,err = pg.fetchCronJobs(ctx,query)
	return
}

func (pg *pgCronJob) DeleteCronJob(ctx context.Context,d r.CronJob)(err error){
	return
}

func (pg *pgCronJob) StartCronJob(ctx context.Context,d r.CronJob)(err error){
	return
}

func(pg *pgCronJob)fetchCronJobs(ctx context.Context,query string,args ...interface{})(res []r.CronJob,err error){
	rows,err := pg.Conn.QueryContext(ctx,query,args...)
	if err != nil {
		return nil,err
	}
	defer func()  {
		rows.Close()
	}()
	res = make([]r.CronJob, 0)
	for rows.Next(){
		t := r.CronJob{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.Expression,
			&t.Tag,
		)
		res = append(res, t)
	}
	return 
}

