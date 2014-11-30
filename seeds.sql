INSERT INTO users (email, password) VALUES
  ('user@example.com', crypt('password', gen_salt('bf')));

INSERT INTO gists (user_id, title, content) VALUES
  (1, 'Some Default Gist', '## Some Markdown');
