-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "student"(
  id serial PRIMARY KEY,
  name varchar(80) NOT NULL,
  birth_day varchar(10),
  display_color varchar(7) NOT NULL DEFAULT '#fefefe',
  picture text,
  gender varchar(6),
  payment_style varchar(7) NOT NULL,
  payment_type varchar(8) NOT NULL,
  payment_type_value float8,
  settlement_style varchar(12) NOT NULL,
  settlement_style_value integer,
  settlement_style_day integer,
  parent_name varchar(80),
  parent_phone_number varchar(20),
  house_address varchar(80),
  house_identifier varchar(80),
  house_coordinate_latitude float8,
  house_coordinate_longitude float8,
  is_deleted boolean NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "student";
-- +goose StatementEnd
