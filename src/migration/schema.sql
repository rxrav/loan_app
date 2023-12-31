create table if not exists loan_users (
    id UUID primary key,
    user_name VARCHAR(256) unique not null,
    first_name VARCHAR(256),
    last_name VARCHAR(256),
    user_age int not null,
    social_number VARCHAR(256) unique not null
);

create table if not exists loan_applications (
	id UUID primary key,
    user_name VARCHAR(256) not null,
    applied_amount int not null,
    approved_amount int default 0,
    status_ varchar(127),
    constraint fk_user_name_application foreign key(user_name) references loan_users (user_name)
);
