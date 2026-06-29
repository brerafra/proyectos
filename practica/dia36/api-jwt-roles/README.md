## API User con middleware auth y con generación de tokens JWT

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
    
4. Login de acceso con Email y Password donde se valida mediante Hash contra el password en base de datos si la validación es correcta, se genera un Token JWT con duración de 15 minutos

5. Middleware JWTAuthorization que valida el Token enviado en Authorization si es vigente y valido.

6. Se utiliza el endpoint protegido por el middleware:
    /user?id= -> para buscar un usuario por id

7. Se utiliza el endpoint protegido por el middleware (con permissions W-> escribir, R-> leer)
    /createuser ->crear un usuario.

