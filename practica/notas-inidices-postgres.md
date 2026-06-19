# Indices

Los inidices son objetos de base de datos, mejoran la velocidad de las operaciones de recuperacion de datos.

## Utilización

optimizan el rendimiento de las consultas SELECT y las clausulas WHERE.

CREATE INDEX index_name ON TABLE (column1, column2, ...)

#### Indice básico:

CREATE INDEX idx_customer_name
ON customers(customer_name);

-> este ejemplo crea un índice en la columna "customer_name"  de la tabla "customers" aceletera las busquedas en esta columna.

