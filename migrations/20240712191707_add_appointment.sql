-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "calendar_day"(
  id serial PRIMARY KEY,
  day integer NOT NULL,
  month integer NOT NULL,
  year integer NOT NULL
  -- UNIQUE (day, month, year)
);

CREATE TABLE IF NOT EXISTS "appointment"(
  id serial PRIMARY KEY,
  id_calendar_day integer NOT NULL,
  id_student integer NOT NULL,
  start_hour varchar(5) NOT NULL,
  duration integer NOT NULL,
  price float8 NOT NULL,
  is_extra bool NOT NULL DEFAULT false,
  is_deleted bool NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (id_calendar_day) REFERENCES "calendar_day"(id) ON DELETE CASCADE,
  FOREIGN KEY (id_student) REFERENCES "student"(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "appointment";
DROP TABLE IF EXISTS "calendar_day";
-- +goose StatementEnd
