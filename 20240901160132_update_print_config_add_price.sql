-- +goose Up
-- +goose StatementBegin
ALTER TABLE print_configs
ADD COLUMN price NUMERIC(10, 2) NOT NULL DEFAULT 0.00;

ALTER TABLE print_configs
ALTER COLUMN quantity TYPE INT USING quantity::integer;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
