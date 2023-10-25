CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE ROLE AS ENUM ('ADMIN', 'STUDENT', 'TEACHER');
 
CREATE TABLE IF NOT EXISTS users (
  "uuid" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "username" VARCHAR(32) UNIQUE,
  "email" VARCHAR UNIQUE,
  "full_name" VARCHAR NOT NULL,
  "password_hash" VARCHAR NOT NULL,
  "role" ROLE NOT NULL
);

INSERT INTO users (uuid, username, email, full_name, password_hash, role)
VALUES (
  '0818f7a7-e8df-44c3-a981-52dd6150f2ab',
  'theadmin',
  'theadmin@mail.com',
  'fulladmin',
  '$2a$08$abHt86mhBpv7frwLLb//4ufNcIjhLzRfxXb6XPG1Qa54iyH4Sxw7W',
  'ADMIN'
)
