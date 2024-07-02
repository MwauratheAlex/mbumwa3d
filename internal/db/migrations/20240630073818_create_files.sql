-- +goose Up
-- +goose StatementBegin

create table files (
  id bigint primary key generated by default as identity,
  user_id bigint,
  inserted_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc'),

  local_path varchar(255),
  file_name varchar(255),
  technology varchar(255),
  color varchar(255),

  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table files;
-- +goose StatementEnd