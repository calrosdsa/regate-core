package pg

import (
	"context"
	r "regate-core/domain/repository"
	"database/sql"
	"log"
	"time"
)

type salaRepo struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) r.SalaRepository{
	return &salaRepo{
		Conn: conn,
	}
}
func (p *salaRepo) DisabledExpiredRooms(ctx context.Context){
	query := `update salas set estado = $1 where horas[1]::timestamp <  $2`
	_,err := p.Conn.ExecContext(ctx,query,r.SalaUnAvailable,time.Now())
	if err != nil {
		log.Println(err)
	}
	log.Println("DELETED")
}

func (p *salaRepo) DeleteUnAvailablesSalas(ctx context.Context){
	query := `delete from salas where estado = $1`
	_,err := p.Conn.ExecContext(ctx,query,r.SalaUnAvailable)
	if err != nil {
		log.Println(err)
	}
	log.Println("DELETED")
}