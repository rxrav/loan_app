CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

insert into loan_users (id, user_name, first_name, last_name, user_age, social_number)
values (uuid_generate_v4(), 'johndoe111111', 'John', 'Doe', 35, 'AAA12001');

insert into loan_users (id, user_name, first_name, last_name, user_age, social_number)
values (uuid_generate_v4(), 'peterparker222222', 'Peter', 'Parker', 16, 'AAA98121');

insert into loan_users (id, user_name, first_name, last_name, user_age, social_number)
values (uuid_generate_v4(), 'brucewayne333333', 'Bruce', 'Wayne', 52, 'AAA65091');

insert into loan_users (id, user_name, first_name, last_name, user_age, social_number)
values (uuid_generate_v4(), 'janedoe444444', 'Jane', 'Doe', 26, 'AAA78034');

insert into loan_users (id, user_name, first_name, last_name, user_age, social_number)
values (uuid_generate_v4(), 'clarkkent555555', 'Clark', 'Kent', 52, 'AAA79811');

INSERT INTO loan_applications (id, user_name, applied_amount, approved_amount, status_ )
VALUES ('1da322d1-84ec-4f23-976d-e3a8512f46a9' :: uuid, 'brucewayne333333', 40000, 21600, 'Approved' );

INSERT INTO loan_applications (id, user_name, applied_amount, approved_amount, status_ )
VALUES('96f347b9-c7a0-49dd-97e0-b7a6cac5ac22' :: uuid, 'johndoe111111', 70000, 33600, 'Approved' );

INSERT INTO loan_applications ( id, user_name, applied_amount, approved_amount, status_ )
VALUES ('83a9a934-af42-4e5b-ba18-5e191d45d594' :: uuid, 'johndoe111111', 20000, 9600, 'Approved');