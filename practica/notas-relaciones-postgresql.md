# Relaciones en postgresql
***

- Las relaciones en SQL estan basadas en la teoría de conjuntos, hay que entender los siguientes conceptos claves.

### Claves primarias y foráneas

Una clave primaria (primary key) se trata de un campo o conjunto de campos que identifican de manera única cada registro en una tabla, que tengan el mismo valor en la columna que se ha desginado como clave primaria.

1.  Unicidad: Cada valor en la columna con clave primaria debe de ser unico, no duplicados.
2. No nulidad: No puden ser nulos.
3. Identificación: la clave primaria identifica de manera exclusiva cada fila en la tabla, eto permite una facil referencia y vinculación  entre tablas en la base de datos.

- Para definirla, en la creación de una tabla ponemos la restricción **PRIMARY KEY** 

CREATE TABLE authors(
    author_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(50)
);

La clave foránea representa una relación entre dos tablas por una conexión entre una columna en una tabla y la clave primaria en otra tabla, se utiliza para garantizar la integridad referencial entre los datos, lo que significa que mantiene la consistencia entre las tablas relacionadas.

1. Relación entre tablas: establece relación entre dos tablas, al vincular una columna en una tabla con la clave primaria de otra tabla.
2. Integridad referencial: la clave foranea garantiza que los valores den la columna de la tabla referencia existan como valores en la clave primaria correspondiente de la talba referenciada.

- Para definirla, usamos **FOREING KEY**

CREATE TABLE books(
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    publication_year INTEGER,
    author_id INTEGER,
    CONSTRAINT fk_author_id FOREIGN KEY (author_id) REFERENCES authors(author_id)
);

Ejemplo de insercción de datos en las tablas authors y books

INSERT INTO authors (first_name, last_name, email) 
VALUES
('John','Doe', 'john.doe@example.com'),
('Jane','Smithb', 'jane.smith@example.com'),
('Michael','Johnson', 'michael.johnson@example.com');

INSERT INTO books (title, publication_year, author_id) 
VALUES
('Introduction to PostgreSQL',2020, 1),
('Advanced SQL techniques',2018, 2),
('Database Design Fundamentals',2019, 3);

### Tipos de relaciones

Se refieren a la manera en que las tablas de una base de datos estan vinculadas entre sí, las cuales son:

* Relación uno a uno.
* Relación uno a muchos.
* Relación muchos a muchos.

#### Relacion uno a uno (1:1)

Significa que una fila de una tabla esta asociada a una fila en otra tabla y viceversa, ejemplo:

CREATE TABLE contries(
    country_id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE capitals(
    capital_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    country_id INT UNIQUE,
    CONSTRAINT fk_country_id FOREIGN KEY (country_id) REFERENCES countries(country_id)
);

y algunos ejemplos de insercción:

INSERT INTO countries (name)
VALUES  ('United States'),
        ('United kingdom'),
        ('France');

INSERT INTO capitals (name, country_id)
VALUES  ('Washington D.C.', 1),
        ('London.', 2),
        ('Paris', 3);

#### Relación uno a muchos (1:N)

significa que una fila en una tabla está asociada con cero o mas filas en otra tabla, pero una fila en la segunda tabla está asociada con exactamente una fila en la primera tabla. 

Ejemplo: un cliente puede tener multiples pedidos pero cada pedido esta asociado a un solo cliente. en este caso la relación seria 1:N.

CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255)
);

CREATE TABLE orders(
    order_id SERIAL PRIMARY KEY,
    order_date DATE,
    total_amount DECIMAL(10, 2),
    customer_id INT,
    CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

y el ejemplo de insercción seria:

INSERT INTO customers(name, email) VALUES
    ('Juan', 'juan@example.com');

INSERT INTO orders (customer_id, order_date, total_amount) VALUES
    (1, '2026-01-14', 150.00),
    (1, '2026-01-15', 200.00),
    (1, '2026-01-16', 100.00);

#### Relación uno a muchos (N:M)

significa que múltiples registros en una tabla están relacionados con múltiples registros en otra tabla y viceversa.

Para implementar una relación muchos a muchos en una db se utiliza una tabla intermedia, tambien conocida como tabla de union o tabla de asociación. Esta tabla contienelas claves primarias de ambás tablas relacionadas y puede incluir atributos adicionales si es necesario.

CREATE TABLE students(
    student_id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE courses(
    course_id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE student_courses(
    student_course_id SERIAL PRIMARY KEY,
    student_id INT,
    course_id INT,
    CONSTRAINT fk_student_id FOREIGN KEY (student_id) REFERENCES students(student_id),
    CONSTRAINT fk_course_id FOREIGN KEY (course_id) REFERENCES courses(course_id),
    CONSTRAINT uc_student_course UNIQUE (student_id, course_id)
);

y aquí algunos ejemplos de insercciones

INSERT INTO students(name) VALUES
    ('juan'),
    ('maria'),
    ('pedro');

INSERT INTO courses(name) VALUES
    ('Matemáticas'),
    ('Historia'),
    ('Ingles');

INSERT INTO student_courses (student_id, course_id) VALUES
    (1,1), -- Juan esta en matematicas
    (1,2), -- Juan esta en historia
    (2,2), -- María esta en historia
    (2,3), -- Maria esta en Ingles
    (3,1), -- pedro esta en matematicas


### Acciones en cascada

Las acciones en cascada, se ejecutan automáticamente cuando ocurre una operación en una tabla principal y sta acción afecta a una tabla secundaria que tiene una relación con la tabla principal a través de una clave foranea.

* ON DELETE CASCADE: Cuando un registro en la tabla principal es eliminado, todos los registros relacionados en la tabla secundaria tambien seran eliminados automáticamente. esto garantiza que no haya registros huérfanos.

CREATE TABLE Invoices (
    invoice_id SERIAL PRIMARY KEY,
    invoice_date DATE
);

CREATE TABLE invoice_items(
    item_id SERIAL PRIMARY KEY,
    invoice_id INT,
    item_name VARCHAR(255),
    CONSTRAINT fk_invoice_id FOREIGN KEY (invoice_id) REFERENCES invoices(invoice_id) ON DELETE CASCADE
);

INSERT INTO invoices (invoice_date) VALUES ('2024-02-13');
INSERT INTO invoice_items (invoice_id, item_name) VALUES
(1, 'Producto A'),
(1, 'Producto B'),
(1, 'Producto C'),

select * FROM invoice_items;
DELETE FROM invoices WHERE invoice_id = 1;
SELECT * FROM  invoice_items;

* ON UPDATE CASCADE: cuando la clave primaria de un registro en la tabla principal es modificada, los valores en la tabla secundaria también se actualizarán automáticamente.

CREATE TABLE departments (
    department_id SERIAL PRIMARY KEY,
    department_name VARCHAR(100)
);

INSERT INTO departments(department_name) VALUES ('ventas');


CREATE TABLE employess (
    employee_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    department_id INT,
    CONSTRAIT fk_department_id FOREIGN KEY (department_id) REFERENCES departments(department_id)
);

INSERT INTO employees (name, department_id) VALUES ('Juan', 1);
INSERT INTO employees (name, department_id) VALUES ('María', 1);

UPDATE departments SET department_id = 100 WHERE department_id = 1;

SELECT * FROM employees;