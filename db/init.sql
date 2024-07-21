CREATE DATABASE chalk_mvp;
\c chalk_mvp;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "CONFERENCE" (
  CODE VARCHAR(20) NOT NULL PRIMARY KEY,
  ADMIN VARCHAR(20) NOT NULL,
  CREATED_AT TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  ACTIVE BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS "STUDENT" (
  ID VARCHAR(255) NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  NAME VARCHAR(255) NOT NULL,
  EMAIL VARCHAR(255) NOT NULL UNIQUE,
  PASSWORD VARCHAR(255) NOT NULL,
  UNIVERSITY VARCHAR(500) NOT NULL,
  YEAR_OF_GRADUATION INTEGER NOT NULL,
  SKILLS VARCHAR(20) ARRAY,
  DESCRIPTION VARCHAR(500)
);