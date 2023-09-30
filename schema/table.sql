CREATE TABLE go_workshop_db.user (
    user_id bigint UNSIGNED not null auto_increment,
    username varchar(200) not null,
    password varchar(200) not null,
    role int not null,
    create_time timestamp not null default CURRENT_TIMESTAMP,
    update_time timestamp not null default CURRENT_TIMESTAMP,
    primary key (user_id),
    unique(username)
);

CREATE TABLE go_workshop_db.user_detail (
    detail_user_id bigint UNSIGNED not null auto_increment,
    user_id bigint not null,
    name varchar(200) not null,
    address varchar(200) not null default '',
    school_name varchar(200) not null default '',
    create_time timestamp not null default CURRENT_TIMESTAMP,
    update_time timestamp not null default CURRENT_TIMESTAMP,
    primary key (detail_user_id),
    unique(user_id)
);