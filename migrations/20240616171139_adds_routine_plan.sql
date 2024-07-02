-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "routine_plan"(
  id serial PRIMARY KEY,
  id_student integer,
  week_day varchar(9) NOT NULL,
  start_hour bigint NOT NULL,
  duration bigint NOT NULL,
  price float8 NOT NULL,
  is_deleted boolean NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (id_student) REFERENCES "student"(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "routine_plan";
-- +goose StatementEnd
