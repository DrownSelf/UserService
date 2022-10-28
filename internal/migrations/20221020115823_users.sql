-- +goose Up
-- +goose StatementBegin
alter table users drop column "rating";
create table if not exists ratings(
    "id" int primary key references users("id"),
    "average_rating" real,
    "rating" real ARRAY
);

create or replace procedure AddRating("input_id" int, "rating_value" real)
    language plpgsql as $$
    begin
        update ratings
            set "rating" = array_append("rating", "rating_value")
        where "id" = "input_id";

        update ratings
            set "average_rating" = (select avg(arr) from unnest("rating") as arr)
        where "id" = "input_id";
    end
    $$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop procedure AddRating(input_id int, rating_value real);
drop table ratings;
alter table users add column "rating" int;
-- +goose StatementEnd