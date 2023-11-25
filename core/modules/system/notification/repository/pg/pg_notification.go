package repository

import (
	"context"
	"database/sql"
	r "regate-core/domain/repository"
	// "time"
)

type pgNotication struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) r.NotificationRepository {
	return &pgNotication{
		Conn: conn,
	}
}



func (p *pgNotication) GetInfoText(ctx context.Context,id int)(res r.InfoText,err error){
	query := `select id,title,content from info_text where id = $1`
	err = p.Conn.QueryRowContext(ctx,query,id).Scan(&res.Id,&res.Title,&res.Content,)
	return
}



// func(pg *pgNotication)fetchCronJobs(ctx context.Context,query string,args ...interface{})(res []r.CronJob,err error){
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

