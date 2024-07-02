-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"(
  id serial PRIMARY KEY,
  name varchar(80) NOT NULL,
  email varchar(80) NOT NULL UNIQUE,
  provider varchar(6) NOT NULL, -- Google
  picture text,
  email_verified boolean NOT NULL DEFAULT false,
  is_deleted boolean NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE "student" ADD COLUMN id_user integer;
ALTER TABLE "student" ADD FOREIGN KEY (id_user) REFERENCES "user"(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "student" DROP COLUMN id_user;
DROP TABLE IF EXISTS "user";
-- +goose StatementEnd
