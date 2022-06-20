drop table if exists orders;
drop table if exists product;
create table if not exists product (
    ID serial primary key,
    Name varchar(20),
    description varchar(100),
    price int
);

drop table if exists customer;
create table if not exists customer (
    id serial primary key,
    name varchar(20),
    phone varchar(10)
);
create table if not exists orders (
    id serial primary key,
    customer_id int,
    product_id int,
    constraint fk_customer_id 
        foreign key (customer_id) references customer(id),
    constraint fk_product_id
        foreign key (product_id) references product(id)
);

insert into product (name,description,price) 
    values ('apple','green apple', 10);
insert into product (name,description,price) 
    values ('orange','green orange', 5);
insert into product (name,description,price) 
    values ('banana','green banana', 8) ;   
insert into customer (name,phone) 
    values('rupak','7847956039');
insert into customer (name,phone)
    values('raj','8618187289');
insert into orders (customer_id,product_id)
    values(1,2);
insert into orders (customer_id,product_id)
    values(2,1);