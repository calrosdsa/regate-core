create table if not exists empresa_settings(
    id serial primary key,
    empresa_id int not null,
    currency_id smallint not null,
    tarifa float4 not NULL,
    CONSTRAINT fk_empresa_setting
    FOREIGN KEY(empresa_id) 
	REFERENCES empresas(empresa_id)
);

insert into empresa_settings(empresa_id,currency_id,tarifa) values(1,1,3.5);

create table if not exists currency (
    id smallserial primary key,
    abb text,
    currency text
);

insert into currency(abb,currency) values('BOB','Boliviano');