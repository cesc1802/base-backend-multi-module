-- +goose Up
create table if not exists users (
    id serial not null,
    login_id varchar(50) not null ,
    "password" varchar(100) not null ,
    salt varchar(50) not null ,
    status varchar(20) not null ,

    primary key (id)
);

-- +goose Down