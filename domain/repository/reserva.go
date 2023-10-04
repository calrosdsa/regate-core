package repository


type ReservaEstado int8

const (
	ReservaValid   ReservaEstado = 0
	ReservaExpired ReservaEstado = 1
	ReservaCancel  ReservaEstado = 2
	ReservaPaid  ReservaEstado = 3
)

type ReservaType int8

const  (
	ReservaTypeApp ReservaType = 1
	ReservaTypeLocal ReservaType = 2
	ReservaTypeSala ReservaType = 3
)
