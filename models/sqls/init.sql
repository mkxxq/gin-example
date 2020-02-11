
-- rule:
--   1. every field must is not null.
--   2. every field must have a default value.

CREATE TABLE test.t_user (
	id BIGINT PRIMARY KEY auto_increment NOT NULL,
	name varchar(100) DEFAULT '' NOT NULL
)ENGINE=InnoDB DEFAULT ChARSET=utf8;
