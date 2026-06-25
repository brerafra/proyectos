# Hash a password using bcrypt

El paquete golang.org/x/crypto/bcrypt convierte las constraseñas en un hash criptográfico
de solo una via, integrando automáticamente  un valor aleatorio(sal) para evitar ataques de diccionario y aplicando un factor de costo para retrasar los ataques de fuerza bruta.


El proceso es el siguiente:

1. Generar el Hash