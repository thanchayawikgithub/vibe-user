create table user (
    id uuid primary key default gen_random_uuid(),
    email varchar(255) not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp
);

create unique index idx_user_email on user (email);
