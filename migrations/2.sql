USE userdb;

CREATE TABLE teacher (
    id INT(6) AUTO_INCREMENT PRIMARY KEY,
    formation VARCHAR(100),
    user_id INT(6),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE student (
     id INT(6) AUTO_INCREMENT PRIMARY KEY,
     description VARCHAR(100),
     user_id INT(6),
     FOREIGN KEY (user_id) REFERENCES users(id)
);