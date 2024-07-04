-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN has_printer BOOLEAN DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN has_printer;
-- +goose StatementEnd
