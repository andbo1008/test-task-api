create table posts(
    id serial primary key not null,
    title varchar(255) not null,
    content varchar(255) not null,
    create_at timestamp not null,
    update_at timestamp not null
);


