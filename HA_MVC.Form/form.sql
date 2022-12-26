create database form;
use form;

CREATE TABLE Form (
    id int,
    fullname varchar(255),
    companyname varchar(255),
    businessphone varchar(255),
    email varchar(255),
	message varchar(255)
);
select * from Form;
drop table Form;