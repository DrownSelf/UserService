-- +goose Up
-- +goose StatementBegin
alter table users drop constraint "users_phoneNumber_email_key";
create unique index "users_phoneNumber_key" on users("phoneNumber");
create unique index "users_email_key" on users("email");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
