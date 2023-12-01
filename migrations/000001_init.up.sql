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
    fio varchar(100) not null
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