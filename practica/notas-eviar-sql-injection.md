# Evitar sql injection usando Golang

Para evintar la inyección de SQL en go, nunca debes concatenar variables o entradas de usuario directamente en tus consultas. En su lugar utiliza siempre sentencias preparadas (prepared statements) o consultas parametrizadas, permitiendo que la base de datos separe el código sql de los datos aportados por el.

1. Consultas Parametrizadas (Manejo estandar)

El paquete database/sql utiliza los simbolos ? o $1 $2, como marcador de posición.

Forma correcta:

query := "SELECT id FROM users WHERE username= ?"
err := db.QueryRow(query, username).Scan(&userID)

//Nunca concqtenar cadenas, esto permite intyeccion SQL
query := fmt.Sprintf("SElect id From users WHERE username='%s'", username)

2. Uso de ORBM (Ejemplo GORM)

Si utilizas bibliotecas de mapeo-objeto-relacion, la regla es la misma utiliza parámetros en lugar de interpolar cadenas directamente.

db.Where("username= ?", username).First(&user)

la mayoria de los ORM's parametrizan las entradas por defecto en sus metodos.

3. Buenas prácticas adicionales

* Limitar privilegios de base de datos: asegurate de que el usuario de la base de datos utilizados por tu aplicacion solo tenga los permisos estrictamente necesarios (por ejemplo, que no puede borrar tablas ni alterar permisos).

* Validación de entradas: Asegúrate de validar los datos del lado del servidor antes de enviarlos a la base de datos (ejemplo longitud, tipo de dato formato del mail)

* Manejo de errores, no expongas los detalels de los errores de la base de datos directamente al usuario final, podrian tener información sobre la estructura de tus tablas o consultas.

