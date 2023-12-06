-- РОЛИ -- РОЛИ -- РОЛИ -- РОЛИ --

create table if not exists roles
(
    id        serial
        constraint roles_pk
            primary key,
    code_name varchar(10) not null
        constraint roles_pk2
            unique,
    title     varchar(50) not null
        constraint roles_pk3
            unique
);
INSERT INTO roles (id, code_name, title) VALUES (1, 'admin', 'Администратор');
INSERT INTO roles (id, code_name, title) VALUES (2, 'trainer', 'Тренер');
INSERT INTO roles (id, code_name, title) VALUES (3, 'member', 'Участник');

-- УЧАСТНИКИ -- УЧАСТНИКИ -- УЧАСТНИКИ -- УЧАСТНИКИ --

create table if not exists members
(
    id  serial
        constraint members_pk
            primary key,
    login    varchar(50)  not null
        constraint members_pk2
            unique,
    password varchar(32)  not null,
    fio varchar(100) not null,
    role_id  integer      not null
        constraint users_roles_id_fk
            references roles
);

-- ТРЕНЕРЫ -- ТРЕНЕРЫ -- ТРЕНЕРЫ -- ТРЕНЕРЫ --

create table if not exists trainers
(
    id        serial
        constraint trainer_pk
            primary key,
    member_id integer not null
        constraint trainer_members_id_fk
            references members
            on delete cascade
);

-- АБОНЕМЕНТЫ -- АБОНЕМЕНТЫ -- АБОНЕМЕНТЫ -- АБОНЕМЕНТЫ --

create table if not exists subscriptions
(
    id       serial
        constraint subscriptions_pk
            primary key,
    title    varchar(50) not null
        constraint subscriptions_pk2
            unique,
    duration integer     not null,
    ruble    integer     not null,
    penny    integer     not null
);

-- УСЛУГИ -- УСЛУГИ -- УСЛУГИ -- УСЛУГИ --

create table if not exists services
(
    id       serial
        constraint service_pk
            primary key,
    title    varchar(50) not null,
    duration integer     not null,
    ruble    integer     not null,
    penny    integer     not null,
    constraint service_pk2
        unique (title, duration)
);

-- АБОНЕМЕНТЫ УЧАСТНИКОВ -- АБОНЕМЕНТЫ УЧАСТНИКОВ -- АБОНЕМЕНТЫ УЧАСТНИКОВ -- АБОНЕМЕНТЫ УЧАСТНИКОВ --

create table if not exists member_subscriptions
(
    id              serial
        constraint member_subscriptions_pk
            primary key,
    member_id       integer                             not null
        constraint member_subscriptions_members_id_fk
            references members,
    subscription_id integer                             not null
        constraint member_subscriptions_subscriptions_id_fk
            references subscriptions,
    created_at      timestamp default CURRENT_TIMESTAMP not null,
    expiration      timestamp                           not null
);