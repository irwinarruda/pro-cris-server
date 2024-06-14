-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS students(
  id serial PRIMARY KEY,
  name varchar(80) NOT NULL,
  birth_day date,
  display_color varchar(7) NOT NULL DEFAULT '#fefefe',
  picture text,
  parent_name varchar(80),
  parent_phone_number varchar(20),
  house_address varchar(80),
  house_identifier varchar(80),
  house_coordinate_latitude float8,
  house_coordinate_longitude float8,
  base_price float8 NOT NULL DEFAULT 0.0,
  is_deleted boolean NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS students;
-- +goose StatementEnd
