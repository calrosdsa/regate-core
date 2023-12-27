create table if not exists cronjob(
    id serial primary key,
    name text,
    expression text,
    tag text
);


insert into cronjob(name,expression,tag) values('Eliminar salas no disponibles','0 0 * * *','delete-unavailables-salas');
insert into cronjob(name,expression,tag) values('Desactivar salas expiradas','0/1 * * * *','disable-expired-rooms');
insert into cronjob(name,expression,tag) values('Crear Deposito del dia','0 0 * * * *','create-deposito');


create table if not exists info_text(
    id serial primary key,
    title text,
    content text
);

insert into info_text(id,title,content)values  (1,'Sucursal estado','{"data":[{"subtitle":"Sucursal Estados","content":"El estado de una sucursal define la disponibilidad de este"}]}');

insert into info_text(id,title,content)values  (2,'Intervalo de tiempo para reservar',
'{
	"data":[
	{
		"subtitle":"",
		"content":"En esta sección de la configuración, podrás definir el intervalo de tiempo en el que los usuarios de la aplicación podrán realizar una reserva, con un mínimo de 30 minutos."
	}
]
}');



