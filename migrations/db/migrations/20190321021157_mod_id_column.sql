
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
alter table todos change id id int(10) unsigned NOT NULL AUTO_INCREMENT;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
alter table todos change id id varchar(255);
