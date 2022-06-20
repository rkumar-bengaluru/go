create table if not exists product (
    ID serial primary key,
    Name varchar(20),
    description varchar(100),
    price int
)

insert into product (name,description,price) 
    values ('apple','green apple', 8.99)
insert into product (name,description,price) 
    values ('orange','green orange', 5)
insert into product (name,description,price) 
    values ('banana','green banana', 8)    