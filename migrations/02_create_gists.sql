CREATE TABLE IF NOT EXISTS gists (
  id serial,
  user_id int NOT NULL,
  title varchar(80) NOT NULL,
  content text,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);
