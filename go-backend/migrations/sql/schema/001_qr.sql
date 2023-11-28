-- +goose Up
CREATE TABLE qr_codes (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL
);

-- +goose Down
DROP TABLE qr_codes;
