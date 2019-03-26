create IF NOT EXISTS schema project_bang;

create table IF NOT EXISTS project_bang.users (
  id bigserial primary key,
  nickname citext unique not null,
  name citext null,
  surname citext null,
  dob date null,
  photo varchar(250) default 'default_img',
  score bigint default 0,
  passwd text not null
);
