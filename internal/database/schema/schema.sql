CREATE TABLE users (
	id text primary key,
	firstname varchar(15) not null,
	lastname varchar(20) not null,
	email varchar(45) not null,
	password text not null,
	role varchar(10) default ''
);


CREATE TABLE borrowers (
    id varchar(40) PRIMARY key not NULL,
    firstname varchar(20) not NULL,
    lastname varchar(20) not null,
    email varchar(40) not null,
    phone varchar(15) default '',
    birth_date varchar(15) default '',
    gender varchar(15) default '',
    nationality varchar(20) default '',
    state_origin varchar(20) default '',
    address varchar(100) default '',
    passport text default '',
    signature text default '',
    job varchar(25) default '',
    job_term smallint default 0,
    income double precision default 0,
    deck text default '',
    has_criminal_record boolean default false,
    offences text[] DEFAULT '{}',
    jail_time smallint default 0,
    kin JSONB DEFAULT '{}', 
    guarantor JSONB DEFAULT '{}',
    nin text default '',
    bvn text default '',
    bank_name text default '',
    account_number text default '',
    identification text default '',
    loan_ids text[] DEFAULT '{}',
    progress smallint default 10,
    credit_score smallint 
);

CREATE TABLE applications (
	id text primary key,
	loanId varchar(15) not null,
   borrowerid text not null,
	term varchar(20) not null,
	type varchar(45) not null,
	amount double precision not null,
	purpose text default '',
   status varchar(15) default 'pending',
   creditworthiness numeric(4,2)
);



ALTER TABLE applications 
ADD FOREIGN KEY(borrowerid) REFERENCES borrowers(id)  ON DELETE CASCADE;