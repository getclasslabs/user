USE userdb;

ALTER TABLE users
    ADD twitter varchar(30),
    ADD facebook varchar(100),
    ADD instagram varchar(30),
    ADD description VARCHAR(100),
    ADD telephone VARCHAR(30),
    ADD address varchar(100);


CREATE TABLE teacher (
    id INT(6) AUTO_INCREMENT PRIMARY KEY,
    formation VARCHAR(100),
    specialization varchar(100),
    working_time int,
    user_id INT(6),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE students (
     id INT(6) AUTO_INCREMENT PRIMARY KEY,
     user_id INT(6),
     FOREIGN KEY (user_id) REFERENCES users(id)
);