
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE users;
CREATE TABLE todos (
    id varchar(255) NOT NULL,
    title varchar(255) NOT NULL,
    text varchar(255) NOT NULL,
    PRIMARY KEY(id)
    // Todo: crated_at・updated_at追加
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
CREATE TABLE users (
    id varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    PRIMARY KEY(id)
    // Todo: crated_at・updated_at追加
);
DROP TABLE todos;
