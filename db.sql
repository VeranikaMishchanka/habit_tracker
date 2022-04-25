create database tracker;

create table users (
user_id serial primary key, 
login varchar(50) not null,
email varchar(100) not null);

create table habits (
habit_id serial primary key,
habit_name varchar(50) not null,
habit_subname varchar(50) check (habit_name is not null));

create table tracking (
record_id serial primary key,
habit_date date not null,
user_id int not null references users (user_id),
habit_id int not null references habits (habit_id));

