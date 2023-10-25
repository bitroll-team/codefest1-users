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
);

CREATE TABLE IF NOT EXISTS followers (
  "uuid" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "followed_user_uuid" UUID,
  "follower_user_uuid" UUID,
  CONSTRAINT fk_followed_users
    FOREIGN KEY ("followed_user_uuid")
      REFERENCES users (uuid),
  CONSTRAINT fk_follower_users
    FOREIGN KEY ("follower_user_uuid")
      REFERENCES users (uuid),
  UNIQUE("followed_user_uuid", "follower_user_uuid")
);
