# Goroutines - Worker pools

El patron worker pool utiliza canales con buffer para poner tareas en cola y canales sin búfer para sincronización, permitiendo a múltiples goroutines procesar trabajos simultáneamente de forma controlada.


### Explicación de conceptos intermedios utilizados.

1. Canales con buffer(make(chan int, numjobs)): El canal jobs almacena los trabajos pendientes. permite que el hilo principal (el productor) deje varias tareas en la cola y avance rápidamente sin bloquearse de inmediato, siempre que no se supere la capacidad del búfer.

2. Dirección del Canal: En la funcion worker, jobs <- chan int indica que el canal es de solo lectura (para consumir trabajos) y results chan<- int es de solo escritura para devolver resputas.

3. close(jobs): Escencial para envitar un error de deadlock. Cierra el canal una vez que se enviaron todas las tareas, lo que permite que el ciclo for job:=range jobs en los workers termine de forma natural y el programa continúe.

4. Sincronización mediante canales: El ciclo final <- results funciona como un punto de encuentro (wait) bloque el hilo principal hasta que reciba una resputa por cada trabajo enviado, garantizando que todoe el procesamiento finalice antes de que el programa termine.
