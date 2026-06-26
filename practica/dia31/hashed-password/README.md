# Hash a password using bcrypt

El paquete golang.org/x/crypto/bcrypt convierte las constraseñas en un hash criptográfico
de solo una via, integrando automáticamente  un valor aleatorio(sal) para evitar ataques de diccionario y aplicando un factor de costo para retrasar los ataques de fuerza bruta.


El proceso es el siguiente:

1. Generar el Hash


Key concepts to keep in mind

Automatic Salting: you don't need to manually generate or store a salt. The "Generate from password" function automatically creates a random salt and embeds it directly into the final return has string.

Cost factor: the second argument is the work factor using cost sets this value to 10 you can increase this integer (up to 31) to make the hasing process slower, which provieds better resistance against brute force attacks as coputer hadware advances.

max length limit: by cryptographic design, the bcrypt algorithm ignores any characters bedyond 72 bytes. if your application expects exceptionally long text entries, consider hashing the input with sha-256 first before passing it to brypt