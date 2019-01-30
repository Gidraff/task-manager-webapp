drop table users;
drop table tasks;
drop table notes;
drop table sessions;

create table users (
  id serial primary key,
  name varchar(255),
  email varchar(255) not null unique,
  password varchar(255) not null,
  created_at timestamp not null
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null
);

create table tasks (
  id serial primary key,
  task_name varchar(255) not null unique,
  user_id integer references users(id),
  created_at timestamp not null
);

create table notes (
  id serial primary key,
  note_description varchar(255) not null unique,
  task_id integer references tasks(id),
  created_at timestamp not null
);
