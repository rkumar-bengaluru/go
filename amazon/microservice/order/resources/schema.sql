create table product_order (
    id serial primary key,
    cid int,
    pid int
);

insert into product_order (id,cid,pid) 
    values (1,1, 1);
insert into product_order (id,cid,pid) 
    values (2,1, 2);
insert into product_order (id,cid,pid) 
    values (3,1, 3);