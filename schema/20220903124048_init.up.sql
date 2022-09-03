BEGIN;

CREATE TABLE IF NOT EXISTS files (
  id serial not null unique,
  path varchar(255) not null unique,
  type varchar(255) not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);

COMMIT;
