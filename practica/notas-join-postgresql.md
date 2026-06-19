# Uso de Joins en SQL (postgresql)

Los joins son una operacion que permine combinar filas de dos o mas tablas basandose en la relación entre ellas.

### tipos:

1. INNER JOIN

este tipo de JOIN devuelve las filas que tienen al menos una coincidencia en ambas tablas basandose en una unión especificada, es decir, toma dos o más tablas y las combina en un solo conjunto de datos basados en una condición de coincidencia. solo incluye filas que tienen un valor coincidente en la columna especificada en la condición de union.

CREATE TABLE products(
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price DECIMAL(10, 2)
);

INSERT INTO products(name, price) VALUES
('Product A', 10.99),
('Product B', 15.99),
('Product C', 20.49);

CREATE TABLE sales(
    sale_id SERIAL PRIMARY KEY,
    product_id INT,
    quantity INT,
    CONSTRAINT fk_product_id FOREIGN KEY(product_id) REFERENCES products(product_id)
);

INSERT INTO sales (product_id, quantity) VALUES
(1,2),
(2,1),
(1,3);

SELECT *
FROM sales s
INNER JOIN products p ON s.product_id = p.product_id; 

2. LEFT JOIN

Combina filas de dos tablas basándose en una condición de unión y devuelte todas las filas de la tabla izquierda y las filas coincidentes de la tabla derecha, si no hay filas coincidentes enla tabla derecha, se devuelven NULL en las columnas correspondientes de la tabla derecha.

En otras palabras, un left JOIN devuelve todas las filas de  la tabla izquierda, junto con las filas coincidentes de la tabla derecha.

CREATE TABLE users (
    id_user SERIAL PRIMARY KEY,
    name VARCHAR(100),
    phone VARCHAR(20)
);

INSERT INTO users(name, phone) VALUES
('User A','555-1234'),
('User B','555-5678'),
('User C','555-9012');

CREATE TABLE gifts(
    id_gift SERIAL PRIMARY KEY,
    id_users INT,
    product VARCHAR(100),
    quantity INT,
    CONSTRAINT fk_user_id FOREIGN KEY(id_user) REFERENCES users(id_user)
);

INSERT INTO gifts (id_user, product, quantity) VALUES
(1,'Product X', 2),
(2,'Product Y', 1),
(1,'Product Z', 3);

SELECT *
FROM users u
LEFT JOIN gifts g ON  u.id_users = g.id_users;

3. RIGHT JOIN

es la contraparte del left join este join regresa lo equivalente a la union entre las dos tablas y todo lo concerniente a la segunda tabla.

CREATE TABLE areas (
    id_area SERIAL PRIMARY KEY,
    name VARCHAR(100)
);

INSERT INTO areas (name) VALUES
('Sales'),
('Marketing'),
('Human Resources');

CREATE TABLE interns(
    id_intern SERIAL PRIMARY KEY,
    name VARCHAR(100),
    id_area INT,
    CONSTRAINT fk_area_id FOREIGN KEY (id_area) REFERENCES areas (id_area)
);

INSERT INTO interns (name, id_area) VALUES
('John',1),
('Mary',2),
('Peter',1),
('Laura',NULL);

SELECT * 
FROM interns i
RIGHT JOIN areas a ON i.id_area = a.id_area; 

---


