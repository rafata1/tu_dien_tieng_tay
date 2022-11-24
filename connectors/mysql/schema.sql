create table users
(
    id       int auto_increment
        primary key,
    username varchar(100) not null,
    email    varchar(255) not null,
    phone    varchar(20)  not null,
    password text         not null,
    constraint users_email_uindex
        unique (email),
    constraint users_phone_uindex
        unique (phone),
    constraint users_username_uindex
        unique (username)
);