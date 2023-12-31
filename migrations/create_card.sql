create table public.card
(
    number     bigint      not null
        constraint card_pk
            primary key,
    user_id    integer   not null
        constraint card_user_account_id_fk
            references public.user_account,
    currency text not null,
    created_at timestamp not null,
    updated_at timestamp not null
);
