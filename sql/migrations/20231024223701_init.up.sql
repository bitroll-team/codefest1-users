-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- CREATE TABLE IF NOT EXISTS users (
  -- "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  -- "username" VARCHAR(32) UNIQUE,
  -- "email" VARCHAR UNIQUE,
  -- "full_name" VARCHAR NOT NULL,
  -- "password_hash" VARCHAR NOT NULL
-- );

CREATE TYPE ROLE AS ENUM ('ADMIN', 'STUDENT', 'TEACHER');
 
CREATE TABLE IF NOT EXISTS users (
  "uuid" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "username" VARCHAR(32) UNIQUE,
  "email" VARCHAR UNIQUE,
  "full_name" VARCHAR NOT NULL,
  "password_hash" VARCHAR NOT NULL,
  "role" ROLE NOT NULL
);
