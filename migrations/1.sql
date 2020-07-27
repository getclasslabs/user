USE userdb;

ALTER TABLE users
    ADD nickname VARCHAR(100) unique,
    ADD register INT,
    ADD gender INT,
    ADD first_name VARCHAR(70),
    ADD last_name VARCHAR(70),
    ADD birthDate DATE;

