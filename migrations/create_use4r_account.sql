create table public.user_account
(
    id              integer      not null
        constraint user_account_pk
            primary key,
    username        text         not null,
    hashed_password integer,
    email           varchar(320) not null
);

