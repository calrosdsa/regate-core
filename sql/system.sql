create table if not exists cronjob(
    id serial primary key,
    name text,
    expression text,
    tag text
);

insert into cronjob(name,expression,tag) values('Eliminar salas no disponibles','0 0 * * *','delete-unavailables-salas');
insert into cronjob(name,expression,tag) values('Desactivar salas expiradas','0/1 * * * *','disable-expired-rooms');
insert into cronjob(name,expression,tag) values('Crear Deposito del dia','0 0 * * * *','create-deposito');