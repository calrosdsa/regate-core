package pg

import (
	"context"
	"database/sql"
	"log"
	r "regate-core/domain/repository"
	"time"
	// "time"
)

type salaRepo struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) r.SalaRepository{
	return &salaRepo{
		Conn: conn,
	}
}
func (p *salaRepo) DisabledExpiredRooms(ctx context.Context)(res []int,err error){
	// loc,_ := time.LoadLocation("")
	loc,_ := time.LoadLocation("America/La_Paz")
	// log.Println(time.Now().In(loc))
	query := `update salas set estado = $1 where horas[1]::timestamp < $2 returning sala_id`
	// select sala_id,estado from salas 
	// where horas[1]::timestamp < current_timestamp
	res,err = p.fetchIds(ctx,query,r.SalaUnAvailable,time.Now().In(loc))
	if err != nil {
		log.Println(err)
	}
	// log.Println("DISABLED ROOM")
	return
}
// SELECT current_timestamp::timestamp AT TIME ZONE 'America/La_Paz';
// SELECT current_timestamp::timestamp without time zone
//    AT TIME ZONE 'America/La_Paz';



func (p *salaRepo) DeleteUnAvailablesSalas(ctx context.Context){
	query := `delete from salas where estado = $1`
	_,err := p.Conn.ExecContext(ctx,query,r.SalaUnAvailable)
	if err != nil {
		log.Println(err)
	}
	log.Println("DELETED ROOM")
}

func (p *salaRepo) fetchIds(ctx context.Context, query string, args ...interface{}) (res []int, err error) {
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