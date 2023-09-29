package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	r "regate-core/domain/repository"
	// "time"
)

type pgAuthRepository struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) r.AuthRepository {
	return &pgAuthRepository{
		Conn: conn,
	}
}
func (p *pgAuthRepository) VerifyEmail(ctx context.Context, userId int, otp int) (err error) {
	// var id int
	// query := `update users set estado = $1 where user_id = $2 and otp = $3 returning user_id`
	// err = p.Conn.QueryRowContext(ctx, query, r.UserEnabled, userId, otp).Scan(&id)
	// if id != userId {
	// 	return errors.New("no pudimos proceder con la verificación. Por favor, verifica que el código sea correcto")
	// }
	return
}

func (p *pgAuthRepository) GetUser(ctx context.Context, userId int) (res r.User, err error) {
	// query := `select u.user_id,u.email,u.estado,u.username,p.profile_photo,p.nombre,p.apellido,p.profile_id,u.otp from users as u
	// 	left join profiles as p on p.user_id = u.user_id
	// where u.user_id = $1`
	// err = p.Conn.QueryRowContext(ctx, query, userId).Scan(&res.UserId, &res.Email, &res.Estado, &res.Username,
	// 	&res.ProfilePhoto, &res.Nombre, &res.Apellido, &res.ProfileId, &res.Otp)
	// if err != nil {
	// 	log.Println("dwkdkasmdk", err)
	// 	return
	// }
	return
}

func (p *pgAuthRepository) Login(ctx context.Context, d *r.LoginRequest) (res r.User, err error) {
	query := `select id,uuid,username,email,rol,created_at,estado,last_login
	 from user_core where email = $1 and password = crypt($2, password);`
	err = p.Conn.QueryRowContext(ctx, query, d.Email, d.Password).Scan(&res.Id,
		&res.Uuid,&res.Username,&res.Email,&res.Rol,&res.CreatedAt,&res.Estado,&res.LastLogin,
	)
	if err != nil {
		log.Println(err)
		err = errors.New("las credenciales introducidas son incorrectas")
		return
	}
	return
}
