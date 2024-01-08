-- +goose Up
-- +goose StatementBegin
CREATE TABLE wastes (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    amount NUMERIC(15,2) DEFAULT 0.00 NOT NULL,
    category_id BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS wastes;
-- +goose StatementEnd
