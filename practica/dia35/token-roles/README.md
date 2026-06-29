## API User con middleware auth, JWT y roles  Admin/user en token

En esta practica se realizan las siguientes actividades:

1. CRUD completo User (V3 con password) estas rutas son libres
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

6. Endpoint login al generar el token, ya genera token con rol

5. Se utiliza el endpoint protegido por el middleware por token con rol de admin/user:
    /user?id= -> para buscar un usuario por id

