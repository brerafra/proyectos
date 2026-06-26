CREATE TABLE employees (
    employee_id SERIAL PRIMARY KEY,
    firstname   VARCHAR(50)     NOT NULL,
    lastname    VARCHAR(50)     NOT NULL,
    access_id   INTEGER         NOT NULL,
    manager_id  INTEGER         NOT NULL,
    email       VARCHAR(254)    UNIQUE NOT NULL,
    employee_status BOOLEAN     NOT NULL DEFAULT FALSE,
    date_or_birth   TIMESTAMP   NOT NULL
);

CREATE TABLE managers(
    manager_id SERIAL PRIMARY KEY,
    firstname   VARCHAR(50)     NOT NULL,
    lastname    VARCHAR(50)     NOT NULL,
    access_id   INTEGER         NOT NULL,
    email       VARCHAR(254)    UNIQUE NOT NULL,
    date_orbirth   TIMESTAMP   NOT NULL
);

ALTER TABLE employees
    ADD FOREIGN KEY (manager_id) REFERENCES managers(manager_id);