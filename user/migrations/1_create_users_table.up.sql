CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
  "id"    uuid DEFAULT uuid_generate_v4(),
  "name"  varchar(255),
  "email" varchar(255),
  PRIMARY KEY("id")
);
