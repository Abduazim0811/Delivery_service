create table if not exists clients(
    id  varchar(255) primary key,
    name varchar(65) not null,
    email varchar(65) unique,
    phone  varchar(65) unique,
    created_at varchar(65) not null
);

create table if not exists locations(
    user_id  varchar(255) not null references clients(id) on delete cascade,
    city varchar(65) not null,
    street varchar(65) not null,
    home_number varchar(65)  not null
);