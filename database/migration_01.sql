create table users
(
    id         int auto_increment,
    email      varchar(255) not null,
    name       varchar(255) not null,
    password   varchar(255) not null,
    type       int          not null,
    created_at datetime     not null,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY (id)
);