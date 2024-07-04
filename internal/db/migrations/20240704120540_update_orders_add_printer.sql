-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
ADD COLUMN printer_id INT,
ADD COLUMN print_status VARCHAR(50),
ADD CONSTRAINT fk_printer
    FOREIGN KEY (printer_id)
    REFERENCES users(id);
-- +goose StatementEnd
