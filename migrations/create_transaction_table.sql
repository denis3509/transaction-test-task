
create table public.transaction
(
    id          integer   not null
        constraint transaction_pk
            primary key,
    user_id     integer   not null
        constraint transaction_user_account_id_fk
            references public.user_account,
    card_number integer      not null
        constraint transaction_card_number_fk
            references public.card,
    type        integer   not null,
    currency    text      not null,
    amount      float8    not null,
    created_at  timestamp not null,
    updated_at  timestamp not null
);

