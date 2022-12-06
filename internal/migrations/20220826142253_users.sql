-- +goose Up
-- +goose StatementBegin
alter table users drop constraint uniquephone;
alter table users drop constraint uniqueemail;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
