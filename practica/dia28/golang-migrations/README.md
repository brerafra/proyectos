# Migraciones en golang

en este ejercicio utilizaremos la herramienta golang-migrate la cual se usa para gestionar migraciones de base de datos en Go, soportando postgreSQL mediante el driver database/postgres.

### instalamos la herramienta CLI

$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey| apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate

- para switchearnos como root es
sudo -i

2. Creamos el archivo de migración usando el siguiente comando.

migrate create -ext sql -dir database/migration/ -seq init_mg

-> Donde:
    -seq    -> genera una versión secuencial
    init_mg -> es el nombre de la migración
3. Se generan dos archivos de los cuales:

up -> este archivo se usa para implementar los cambios deseados a la base de datos.
down -> este archivo deshace los cambios y devuelve la base de datos a su estado previo.

4. Formato de los archivos SQL son

{version}_{titulo}.down.sql
{version}_{titulo}.up.sql

5. Poner los cambios requeridos en cada migración.
(ejemplo realizado en el archivo up de este ejemplo)

6. escribir el siguiente comando para realizar la migración

migrate -path /migrations -database "postgresql://username:password@localhost:port/database_name?sslmode=disable" -verbose up

ejemplo real:

migrate -path /migrations -database "postgresql://username:postgres://admin:brerafra@localhost:6432/test?sslmode=disable" -verbose up

migrate -database "YOUR_DATABASE_URL" -path PATH_TO_MIGRATIONS force 1   