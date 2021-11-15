CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS auth_user (
  "id"          uuid DEFAULT uuid_generate_v4(),
  "email"       varchar(255),
  "password"    varchar(255),
  PRIMARY KEY("id")
);
