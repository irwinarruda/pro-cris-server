-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "account"(
  id serial PRIMARY KEY,
  name varchar(80) NOT NULL,
  email varchar(80) NOT NULL UNIQUE,
  provider varchar(6) NOT NULL,
  picture text,
  email_verified boolean NOT NULL DEFAULT false,
  is_deleted boolean NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE "student" ADD COLUMN id_account integer;
ALTER TABLE "student" ADD FOREIGN KEY (id_account) REFERENCES "account"(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "student" DROP COLUMN id_account;
DROP TABLE IF EXISTS "account";
-- +goose StatementEnd
