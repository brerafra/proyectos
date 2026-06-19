# Transacciones postgresql

Una transacción empaqueta varios pasos en una operación, de forma que se completen todos o ninguno.

En el caso de que ocurra algún fallo que impida que se complete la transacción, niguno de los pasos se ejecutan y no afectan a los objdetos de la base de datos.

Comandos ->

BEGIN

Este comando permite que se ejecuten todas las sentencias SQL que necesitamos y las registra en un fichero. ejemplo:

BEGIN;

UPDATE cuentas SET balance = balcance - 100.00 WHERE n_cuenta = 012834;

UPDATE cuentas SET balance = balcance - 100.00 WHERE n_cuenta = 012817;

COMMIT

Este comando se usa para confirmar que todas las sentencias son correctas, por ejemplo si cerramos la conexión antes de ejecutar este comando, no se vera afectada ninguna de las relaciones de la base de datos:

ejemplo:

BEGIN;

INSERT INTO cuentas (n_cuenta, nombre, balance) VALUES(0679283, 'Pepe' 200);

UPDATE cuentas SET balance = balance - 137.00 WHERE  nombre = 'Pepe';

UPDATE cuentas SET balance = balance - 137.00 WHERE  nombre = 'Juan';

SELECT nombre, balance FROM cuentas WHERE nombre = 'Pepe' AND nombre='Juan';

COMMIT;

Comando RollBack

Con este comando podemos desechar las transacciones que se hayan ejecutado, despues de haber realizado y confirmado una transacción, PostgreSQL nos permite anular dicha transacción de forma que no se modifique los datos de nuestra base de datos.

ejemplo:

BEGIN;

"SENTENCIAS SQL"
COMMIT;

ROLLBACK;

-> Para poder usar estos comandos es necesario desactivar el autocommit que vienen en todos los clientes postgresql
