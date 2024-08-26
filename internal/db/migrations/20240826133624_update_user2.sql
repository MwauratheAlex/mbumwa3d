-- +goose Up
-- +goose StatementBegin
ALTER TABLE users 
    ADD COLUMN name VARCHAR,
    ADD COLUMN photo_url VARCHAR;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users 
    DROP COLUMN name,
    DROP COLUMN photo_url;
-- +goose StatementEnd
