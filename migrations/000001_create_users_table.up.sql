CREATE TABLE IF NOT EXISTS users
(
  id SERIAL PRIMARY KEY,
  name varchar(100) not null,
  username varchar(100) not null,
  password varchar(255) NOT NULL
);
