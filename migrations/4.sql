USE userdb;

CREATE TABLE reviews (
    id INT(6) AUTO_INCREMENT PRIMARY KEY,
    comment VARCHAR(255),
    value FLOAT,
    student_id INT(6),
    teacher_id INT(6),
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (teacher_id) REFERENCES teacher(id)
);
