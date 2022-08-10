-- +goose Up
-- +goose StatementBegin
create table users(
    "id" serial primary key,
    "name" text,
    "phoneNumber" text,
    "email" text,
    "password" text,
unique("phoneNumber", "email"));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd