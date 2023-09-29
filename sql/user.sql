create table if not exists user_core(
    id serial primary key,
    uuid uuid unique DEFAULT uuid_generate_v4(),
    username VARCHAR UNIQUE not null,
    password VARCHAR NOT NULL,
    email VARCHAR UNIQUE not null,
    phone VARCHAR,
    rol int DEFAULT 0,
    created_at TIMESTAMP DEFAULT now(),
    estado int DEFAULT 0,
    last_login TIMESTAMP
);

 insert into user_core (username,password,email,phone) values ('admin',crypt('12ab34cd56ef', gen_salt('bf')),'admin@gmail.com','741390560');