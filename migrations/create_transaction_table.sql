
create table public.transaction
(
    id          serial  primary key not null,
    card_number bigint      not null
        constraint transaction_card_number_fk
            references public.card,
    type        integer   not null,
    status        integer   not null,
    currency    text      not null,
    amount      float8    not null,
    created_at  timestamp not null,
    updated_at  timestamp not null
);

