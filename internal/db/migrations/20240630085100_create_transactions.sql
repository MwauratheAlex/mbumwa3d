-- +goose Up
-- +goose StatementBegin

CREATE TABLE transactions (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    file_id BIGINT NOT NULL,
    inserted_at TIMESTAMP(0) NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at TIMESTAMP(0) NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),

    build_time BIGINT NOT NULL,
    quantity VARCHAR(255) NOT NULL,
    price FLOAT8 NOT NULL,
    phone VARCHAR(255) NOT NULL,
    payment_complete BOOLEAN NOT NULL,
    status VARCHAR(255) NOT NULL,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_file FOREIGN KEY (file_id) REFERENCES files(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE transactions;

-- +goose StatementEnd
