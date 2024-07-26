CREATE DATABASE chalk_mvp;
\c chalk_mvp;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "STUDENT" (
  ID VARCHAR(255) NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  NAME VARCHAR(255) NOT NULL,
  EMAIL VARCHAR(255) NOT NULL UNIQUE,
  PASSWORD VARCHAR(255) NOT NULL,
  UNIVERSITY VARCHAR(500),
  YEAR_OF_GRADUATION INTEGER,
  SKILLS VARCHAR(20) ARRAY,
  DESCRIPTION VARCHAR(500),
  DEGREE VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS "BOOKMARKS" (
  EMAIL VARCHAR(255) NOT NULL PRIMARY KEY REFERENCES public."STUDENT"(EMAIL), 
  STUDENT_EMAILS VARCHAR(255) ARRAY
);

CREATE TABLE IF NOT EXISTS "CHAT_CODE" (
  ID1 VARCHAR(255) NOT NULL REFERENCES public."STUDENT"(ID),
  ID2 VARCHAR(255) NOT NULL REFERENCES public."STUDENT"(ID),
  CODE VARCHAR(10) NOT NULL UNIQUE PRIMARY KEY,
);

CREATE TABLE IF NOT EXISTS "CHAT" (
  ID VARCHAR(255) NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  SENDER_ID VARCHAR(255) NOT NULL REFERENCES public."STUDENT"(ID),
  CHAT_CODE VARCHAR(255) NOT NULL REFERENCES public."CHAT_CODE"(CODE),
  MESSAGE TEXT NOT NULL,
  SENT_AT TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
