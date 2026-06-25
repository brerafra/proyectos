docker run --name postgres_db \
    -e POSTGRES_USER=admin \
    -e POSTGRES_PASSWORD=brerafra \
    -e POSTGRES_DB=test \
    -p 6432:5432 \
    -d postgres


para ingresar por consola 

psql -h localhost -p 6432 -U admin -d test
