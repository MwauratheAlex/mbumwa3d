-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders RENAME COLUMN printconfigid TO print_config_id;
ALTER TABLE orders ADD COLUMN checkout_request_id VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders DROP COLUMN checkout_request_id;
ALTER TABLE orders RENAME COLUMN print_config_id TO printconfigid;
-- +goose StatementEnd
