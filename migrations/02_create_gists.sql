CREATE TABLE IF NOT EXISTS gists (
  id serial,
  user_id int,
  title varchar(80) NOT NULL,
  content text
);
