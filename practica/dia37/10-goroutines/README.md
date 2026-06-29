# Elementos claves de las goroutines

* Sync.WaitGroup -> Es la herramienta estándar para coordinar goroutines. Nos permite bloquear la ejecución del programa principal hasta que todas las tareas concurrente hayan notificado de su término.

* defer wg.Done() -> Se asegura de notificar al WaitGroup que la goroutine ha terminado. Justo antes de salir de la funcion anonima.

* go func() -> la palabra clave go es la que convierte la llamada a la funcion goroutine independiente. pasamos la variable i como argumento para evitar problemas de concurrencia con la variable del buckle.