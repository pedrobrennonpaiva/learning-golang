create database if not exists golang_course;

use golang_course;

drop table if exists users;
create table users(
    id int not null auto_increment,
    name varchar(50) not null,
    nickname varchar(50) not null,
    email varchar(50) not null,
    password varchar(200) not null,
    created_at timestamp default current_timestamp() not null,
    primary key (id),
    unique key (email),
    unique key (nickname)
) ENGINE=INNODB;

drop table if exists followers;
create table followers(
    user_id int not null,
    follower_id int not null,
    primary key (user_id, follower_id),
    foreign key (user_id) references users(id) on delete cascade,
    foreign key (follower_id) references users(id) on delete cascade
) ENGINE=INNODB;

drop table if exists posts;
create table posts(
    id int not null auto_increment,
    title varchar(50) not null,
    content varchar(300) not null,
    author_id int not null,
    created_at timestamp default current_timestamp() not null,
    primary key (id),
    foreign key (author_id) references users(id) on delete cascade
) ENGINE=INNODB;

drop table if exists likes;
create table likes(
    id int not null auto_increment,
    user_id int not null,
    post_id int not null,
    primary key (id),
    foreign key (user_id) references users(id) on delete cascade,
    foreign key (post_id) references posts(id) on delete cascade
) ENGINE=INNODB;