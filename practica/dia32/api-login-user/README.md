# API User Con login web

En esta practica se realizan las siguientes actividades:

1. CRUD completo User (V3 con password)
    - Create
    - GetById
    - GetAll
    - GetByEmail
    - Update
    - Delete

2. Se usa el algoritmo Hash para encriptar el password

3. Se usa Postgres
    - Se utiliza postgres Dockerizado
    - Se utiliza la migración dentro de connection
    
4. Se crean las rutas para despachar paginas web:
    - index
    - internal
    - logout
    - control de errores básico

5. Se inicia la utilización de cookies para almacenar el email en una sessión de login y logut.