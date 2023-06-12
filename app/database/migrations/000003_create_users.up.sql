CREATE TABLE users (
  id       BIGSERIAL   primary key,
  email    VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL
)
