-- +goose Up
-- +goose StatementBegin
CREATE TABLE article(
	id serial PRIMARY KEY,
	title VARCHAR(64) NOT NULL,
	content VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS article
-- +goose StatementEnd
