-- CREATE DATABASE app_template ENCODING 'UTF8';

CREATE TABLE IF NOT EXISTS "app_template"(
    id bigint not null primary key,
    app_name varchar(255) not null,
    project_id int not null,
    desc varchar(1024) not null,
    icon varchar(255),
    manifest_loc text,
    status int not null,
    create_time timestamp default CURRENT_TIMESTAMP,
    update_time timestamp
);