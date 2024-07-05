-- +goose Up
-- +goose StatementBegin
-- +goose StatementEnd

ALTER TABLE orders ALTER COLUMN printer_id DROP NOT NULL;
ALTER TABLE orders DROP CONSTRAINT fk_printer;

ALTER TABLE orders ADD CONSTRAINT fk_printer
FOREIGN KEY (printer_id) REFERENCES users(id) ON DELETE SET NULL;
-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
