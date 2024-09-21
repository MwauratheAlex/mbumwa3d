-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
ADD COLUMN printConfigID BIGINT;

ALTER TABLE orders
ADD CONSTRAINT fk_print_config
FOREIGN KEY (printConfigID) REFERENCES print_configs(id);

ALTER TABLE orders
DROP COLUMN file_id,
DROP COLUMN quantity,
DROP COLUMN print_status;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
DROP CONSTRAINT fk_print_config;

ALTER TABLE orders
DROP COLUMN printConfigID;

-- Add the dropped columns back to the `orders` table
ALTER TABLE orders
ADD COLUMN file_id VARCHAR(255),
ADD COLUMN quantity VARCHAR(255),
ADD COLUMN print_status VARCHAR(255);
-- +goose StatementEnd
