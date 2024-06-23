create table if not exists drivers(
    id varchar(255) primary key,
    name varchar(65) not null,
    email varchar(65) unique,
    phone varchar(65) unique,
    working_place varchar(65) not null,
    active varchar(65) check(active in('active', 'inactive'))
);

create table if not exists locations (
    driver_id varchar(255) references drivers(id) on delete cascade ,
    city  varchar(65) not null,
    street varchar(65) not null,
    home_number varchar(65) unique
);

create table if not exists statuses (
    driver_id varchar(255) references drivers(id) on delete cascade,
    vehicle varchar(65) not null,
    join_date varchar(65) not null,
    rating float default 1 check(rating between 0 and 5)
);

