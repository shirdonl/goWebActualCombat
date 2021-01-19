create table if not exists user
(
  uid bigint auto_increment
    primary key,
  name varchar(20) default '' null,
  phone varchar(20) default '' null
)
  charset=utf8mb4
;

