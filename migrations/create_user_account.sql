create table public.user_account
(
    id              serial  primary key not null,
    username        text         not null,
    hashed_password text         not null,
    email           varchar(320) not null
);

