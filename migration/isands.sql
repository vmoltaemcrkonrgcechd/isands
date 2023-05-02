drop database if exists isands;

create database isands;

\connect isands

create table country(
                        country_id uuid default gen_random_uuid(),
                        name varchar(40) not null,
                        primary key (country_id)
);

create table supplier(
                         supplier_id uuid default gen_random_uuid(),
                         name varchar(40) not null,
                         primary key (supplier_id)
);

create table category(
                         category_id uuid default gen_random_uuid(),
                         name varchar(40) not null,
                         primary key (category_id)
);

create table color(
                      color_id uuid default gen_random_uuid(),
                      name varchar(40) not null,
                      primary key (color_id)
);

create table product(
                        product_id uuid default gen_random_uuid(),
                        name varchar(40) not null,
                        price real not null,
                        country_id uuid not null,
                        supplier_id uuid not null,
                        category_id uuid not null,
                        color_id uuid not null,
                        online boolean default false,
                        installment_plan boolean default false,
                        in_stock boolean default false,
                        primary key (product_id),
                        foreign key (country_id)
                            references country(country_id)
                            on delete cascade,
                        foreign key (supplier_id)
                            references supplier(supplier_id)
                            on delete cascade,
                        foreign key (category_id)
                            references category(category_id)
                            on delete cascade,
                        foreign key (color_id)
                            references color(color_id)
                            on delete cascade
);

create table detail(
                       detail_id uuid default gen_random_uuid(),
                       name varchar(40) not null,
                       primary key (detail_id)
);

create table product_detail(
                               product_id uuid,
                               detail_id uuid,
                               value text not null,
                               foreign key (product_id)
                                   references product(product_id)
                                   on delete cascade,
                               foreign key (detail_id)
                                   references detail(detail_id)
                                   on delete cascade
);
