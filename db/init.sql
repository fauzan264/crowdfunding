CREATE DATABASE bwastartup;

USE bwastartup;

CREATE TABLE users(
    id varchar(36) primary key not null,
    name varchar(100) not null,
    occupation varchar(100) not null,
    email varchar(100) not null unique,
    password varchar(255) not null,
    avatar_file_name varchar(255) not null,
    role varchar(255) not null,
    created_by varchar(36) not null default '00000000-0000-0000-0000-000000000000',
    created_at datetime not null default now(),
    updated_by varchar(36) not null default '00000000-0000-0000-0000-000000000000',
    updated_at datetime not null default now()
);

CREATE TABLE campaigns(
    id varchar(36) not null primary key,
    user_id varchar(36) not null,
    title varchar(255) not null,
    short_description varchar(100) not null,
    description text,
    goal_amount int UNSIGNED,
    current_amount int UNSIGNED,
    perks text,
    becker_count int,
    slug varchar(255),
    created_by varchar(36) not null default '00000000-0000-0000-0000-000000000000',
    created_at datetime not null default now(),
    updated_by varchar(36) not null default '00000000-0000-0000-0000-000000000000',
    updated_at datetime not null default now()
);

CREATE TABLE campaign_images(
    id varchar(36) not null primary key,
    campaign_id varchar(36) not null,
    file_name varchar(255) not null,
    is_primary tinyint unsigned not null,
    created_by varchar(36) not null default '00000000-0000-0000-0000-000000000000',
    created_at datetime not null default now(),
    updated_by varchar(36) not null default '00000000-0000-0000-0000-000000000000',
    updated_at datetime not null default now()
);

CREATE TABLE transactions(
    id varchar(36) not null,
    user_id varchar(36) not null,
    campaign_id varchar(36) not null,
    amount int unsigned not null,
    status varchar(50) not null,
    code varchar(50) not null,
    created_by varchar(36) not null default '00000000-0000-0000-0000-000000000000',
    created_at datetime not null default now(),
    updated_by varchar(36) not null default '00000000-0000-0000-0000-000000000000',
    updated_at datetime not null default now()
);

CREATE TABLE status_transaction(
    name varchar(50) PRIMARY KEY
);

INSERT INTO status_transaction VALUES
('pending'),
('process'),
('completed');

ALTER TABLE campaigns ADD CONSTRAINT fk_campaign_user FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE campaign_images ADD CONSTRAINT fk_campaign_images_campaign FOREIGN KEY (campaign_id) REFERENCES campaigns(id);

ALTER TABLE transactions ADD CONSTRAINT fk_transactions_user FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE transactions ADD CONSTRAINT fk_transactions_campaign FOREIGN KEY (campaign_id) REFERENCES campaigns(id);

ALTER TABLE transactions ADD CONSTRAINT fk_transactions_status FOREIGN KEY (status) REFERENCES status_transaction(name);