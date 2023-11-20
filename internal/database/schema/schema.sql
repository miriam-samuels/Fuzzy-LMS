CREATE TABLE users (
	id text primary key,
	firstname varchar(15) not null,
	lastname varchar(20) not null,
	email varchar(45) not null,
	password text not null,
	role varchar(6) default ''
);


CREATE TABLE borrowers (
    id VARCHAR(40) PRIMARY key not NULL,
    firstname VARCHAR(20) not NULL,
    lastname VARCHAR(20) not null,
    email VARCHAR(40) not null,
    phone VARCHAR(15) default '',
    birth_date VARCHAR(15) default '',
    gender VARCHAR(15) default '',
    nationality VARCHAR(20) default '',
    state_origin VARCHAR(20) default '',
    address VARCHAR(100) default '',
    passport text default '',
    signature text default '',
    job VARCHAR(25) default '',
    job_term smallint default 0,
    income VARCHAR(25) default '',
    deck text default '',
    has_criminal_record BOOLEAN default false,
    offences TEXT[] DEFAULT '{}',
    jail_time smallint default 0,
    has_collateral BOOLEAN default false,
    collateral TEXT[] DEFAULT '{}',
    collateral_docs text default '',
    kin JSONB DEFAULT '{}', 
    guarantor JSONB DEFAULT '{}',
    nin text default '',
    bvn text default '',
    bank_name text default '',
    account_number text default '',
    identification text default '',
    loan_ids TEXT[] DEFAULT '{}'
);

CREATE TABLE applications (
	id text primary key,
	loanId varchar(15) not null,
   borrowerid text not null,
	term varchar(20) not null,
	type varchar(45) not null,
	amount float not null,
	purpose text default '',
   status varchar(15) default 'pending',
   creditworthiness float 
);

ALTER TABLE applications 
ADD FOREIGN KEY(borrowerid) REFERENCES borrowers(id)  ON DELETE CASCADE;