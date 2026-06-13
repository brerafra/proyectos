# Tutorial Postgresql

### Comandos bash

#### Ingresar por bash

sudo -u postgres psql

#### Crear uri

postgresql://usuario:contraseña@host:port/nombre_bd

#### Comandos:

\l -> Listar DB's
q -> salir de lista DB's
\c dbname -> Conectar a una db especifica
\df -> listar todas las tablas en la db actual
\du -> lista todos los roles y usuarios 
\d tablename -> describe la estructura de la tabla

***

### Sintaxis SQL

***

#### Crear una tabla:
CREATE TABLE users (id SERIAL PRIMARY KEY, name VARCHAR(100), email TEXT UNIQUE NOT NULL);

***

#### Insertar un registro:
INSERT INTO users (name, email) VALUES ('Alice', 'alice@example');

***

#### Query records
SELECT * FROM users;

***

#### update record
UPDATE users SET name='Alice Smith' WHERE id=1;

***

#### delete a record
DELETE FROM users WHERE id=1;

***

#### drop table enterely
DROP TABLE users;

***

#### To backup a database from the shell (outside psql):
pg_dump -U appuser -d appdb < appdb_backup.sql

***

#### To restore from a backup:
psql -U appuser -d appdb < appdb_backup.sql

***
***

### TroubleShooting Common issues

### verificar que el servicio este corriendo
sudo systemctl start postgresql

sudo systemctl status postgresql

### Verificar que estemos conectado al host y puerto correcto, si queremos permitir que se conecte de un pc externo

sudo nano /etc/postgresql/16/main/postgresql.conf

##### Find the line

#listen_addresses = 'localhost'

##### Change to 

listen_addresses = "*"

### editamos pg_hba.conf para permitir conexiones desde un rango de ip valido

sudo nano /etc/postgresql/16/main/pg_hba.conf

### Añadir la siguiente línea para añadir conexiones con la auteticación del password

host    all     all     192.168.1.0/24  scram-sha-256

### si estas usando firewall para habilitar el puerto:

sudo ufw allow 5432/tcp

