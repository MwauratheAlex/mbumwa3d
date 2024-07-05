-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
    DROP CONSTRAINT IF EXISTS fk_printer;

ALTER TABLE orders
    DROP COLUMN IF EXISTS printer_id;

ALTER TABLE orders
    ADD COLUMN printer_id  BIGINT NULL;

ALTER TABLE orders
    ADD CONSTRAINT fk_printer
        FOREIGN KEY (printer_id)
        REFERENCES users(id)
        ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
    DROP CONSTRAINT IF EXISTS fk_printer;

ALTER TABLE orders
    ALTER COLUMN printer_id TYPE INT,
    ALTER COLUMN printer_id SET NOT NULL;

ALTER TABLE orders
    ADD CONSTRAINT fk_printer
        FOREIGN KEY (printer_id)
        REFERENCES users(id);
-- +goose StatementEnd
