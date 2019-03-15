create schema project_bang;

create table project_bang.users (
  id bigserial primary key,
  nickname citext unique not null,
  name citext null,
  surname citext null,
  dob date null,
  photo varchar(250) default 'default_img',
  score bigint default 0,
  passwd text not null
);

insert into project_bang.users (nickname, name, surname, dob, passwd)
    values ('bob', 'bob', 'bob', '1971-07-13', 'qwerty');