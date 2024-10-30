-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "settlement"(
  id serial PRIMARY KEY,
  id_account integer NOT NULL,
  id_student integer NOT NULL,
  payment_style varchar(7) NOT NULL,
  payment_type varchar(8) NOT NULL,
  payment_type_value float8,
  settlement_style varchar(12) NOT NULL,
  settlement_style_value integer,
  settlement_style_day integer,
  start_date timestamp NOT NULL,
  end_date timestamp NOT NULL,
  is_settled bool NOT NULL DEFAULT false,
  is_deleted bool NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (id_account) REFERENCES "account"(id) ON DELETE CASCADE,
  FOREIGN KEY (id_student) REFERENCES "student"(id) ON DELETE CASCADE
);

ALTER TABLE "appointment" ADD COLUMN id_settlement integer;
ALTER TABLE "appointment" ADD FOREIGN KEY (id_settlement) REFERENCES "settlement"(id) ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "appointment" DROP COLUMN id_settlement;
DROP TABLE IF EXISTS "settlement";
-- +goose StatementEnd
