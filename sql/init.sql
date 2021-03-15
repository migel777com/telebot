CREATE TABLE users(chat_id serial not null unique, email varchar(65) not null,  verified boolean not null, vkey text not null, create_time timestamp);
CREATE TABLE books(book_id serial not null unique, course_id integer not null,  subject varchar(65) not null,  book_name varchar(65) not null, book_link text not null);
INSERT INTO books(course_id, subject, book_name, book_link) VALUES (1, 'ICT', 'ICT Basics. How to talk to PC', 'https://drive.google.com/uc?export=download&id=14MKP04DQFIGgctijW1JxPC-DNUynOkv5');
INSERT INTO books(course_id, subject, book_name, book_link) VALUES (1, 'MATH', 'Discrete Mathematics and Its Applications', 'https://drive.google.com/uc?export=download&id=1FOdcpnv_V_UzR_HEvopykJFNCEkWK7p0');
INSERT INTO books(course_id, subject, book_name, book_link) VALUES (2, 'MATH', 'A Collection of Problems on Course of Mathematical Analysis', 'https://drive.google.com/uc?export=download&id=1-yBffqKWRBASQEK9HVwNZihNRE_j2GVZ');