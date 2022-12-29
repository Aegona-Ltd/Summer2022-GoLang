create database form;
use form;

CREATE TABLE Form(
  id int NOT NULL AUTO_INCREMENT,
  fullname varchar(150) CHARACTER SET utf8 NOT NULL,
  companyname varchar(150) CHARACTER SET utf8 NOT NULL,
  businessphone varchar(150) CHARACTER SET utf8 NOT NULL,
  email varchar(150) CHARACTER SET utf8 NOT NULL,
  message varchar(150) CHARACTER SET utf8 NOT NULL,
  createtime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatetime timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
select * from Form;
drop table Form;