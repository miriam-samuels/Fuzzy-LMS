# create the database

CREATE TABLE public.users (
	id text primary key,
	firstname varchar(15) not null,
	lastname varchar(20) not null,
	email varchar(45) not null,
	password text not null,
	role varchar(6) default ''
);
