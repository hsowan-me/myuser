drop database if exists `my_user`;
create database my_user;
use my_user;

drop table if exists `user`;
create table `user` (
    `id` bigint primary key auto_increment,
    `create_time` timestamp default current_timestamp,
    `update_time` timestamp default current_timestamp on update current_timestamp,
    `username` varchar(30) unique,
    `password` varchar(30) not null,
    `salt` char(32),
    `phone` char(11) unique,
    `email` varchar(50) unique,
    `avatar` varchar(200),
    `role` tinyint(1) comment '角色, 0admin, 1student, 2company'
);

drop table if exists `student`;
create table `student` (
    `id` bigint primary key auto_increment,
    `create_time` timestamp default current_timestamp,
    `update_time` timestamp default current_timestamp on update current_timestamp,
    `user_id` bigint,
    `name` varchar(30),
    `school` varchar(50),
    `major` varchar(50),
    `gender` bit,
    `graduate_year` smallint,
    `education` varchar(10) comment '本科 硕士 博士',
    foreign key (`user_id`) references `user`(`id`)
);

drop table if exists `company`;
create table `company` (
    `id` bigint primary key auto_increment,
    `create_time` timestamp default current_timestamp,
    `update_time` timestamp default current_timestamp on update current_timestamp,
    `user_id` bigint,
    `short_name` varchar(30) comment '企业简称',
    `full_name` varchar(50) comment '企业全称',
    `logo` varchar(200),
    `address` varchar(100),
    `trade` varchar(50) comment '行业',
    `introduce` varchar(200),
    foreign key (`user_id`) references `user`(`id`)
);