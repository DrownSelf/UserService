-- +goose Up
-- +goose StatementBegin
drop index users_email_key;
drop index users_phoneNumber_key;
alter table users add column is_deleted boolean not null;
alter table users add constraint uniquePhone unique ("phoneNumber");
alter table users add constraint uniqueEmail unique ("email");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
