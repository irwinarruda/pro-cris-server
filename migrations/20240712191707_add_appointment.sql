-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "appointment"(
  id serial PRIMARY KEY,
  id_account integer NOT NULL,
  id_student integer NOT NULL,
  start_hour varchar(5) NOT NULL,
  calendar_day timestamp NOT NULL,
  duration integer NOT NULL,
  price float8 NOT NULL,
  is_extra bool NOT NULL DEFAULT false,
  is_paid bool NOT NULL DEFAULT false,
  is_deleted bool NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (id_student) REFERENCES "student"(id) ON DELETE CASCADE,
  FOREIGN KEY (id_account) REFERENCES "account"(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "appointment";
-- +goose StatementEnd
