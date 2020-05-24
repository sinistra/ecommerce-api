create table if not exists items
(
    id int auto_increment
        primary key,
    code varchar(30) not null,
    title varchar(50) not null,
    description varchar(200) not null,
    seller int null,
    picture varchar(100) not null,
    price decimal(10,2) default 0.00 null,
    qty_avail int not null,
    qty_sold int not null,
    status varchar(10) not null,
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

create table if not exists users
(
    id int auto_increment
        primary key,
    first_name varchar(30) not null,
    last_name varchar(30) not null,
    email varchar(100) not null,
    status varchar(10) default 'unverified' not null,
    password varchar(100) not null,
    uuid varchar(36) null,
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint users_email_uindex
        unique (email),
    constraint users_uuid_uindex
        unique (uuid)
);
