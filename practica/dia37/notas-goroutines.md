# Goroutines

Son hilos de ejecución ligeros gestionados por el entorno de ejecución de Go. permiten manejar concurrencia de forma masiva y eficiente dentro de una aplicación. Es la herramienta principal de Go para ejecutar multiples tareas simultaneas sin consumer altos recursos de los hilos tradicionales de sistemas operativos.

## Caracteristicas principales

#### Extremadamente ligeras

Comienzan ocupando una cantidad minima de memoria (~2KB) y tienen pilas expandibles, lo que permite crear miles de ellas sin agotar la memoria del sistema.

#### Gestion eficiente

No son hilos del sistema, son gestionadas por el runtime de Go, el cual las distribuye y balancea (multiplexea).

#### Facil uso

Crear una Goroutine es muy sencillo solo requiere anteponer la palabra reservada go antes de cualquier función.


## Diferencias con los hilos tradicionales

A diferencia de los hilos tradicionales de hardware o del SO, las goroutines no requieren costosos cambios de contexto ni llamadas al espacio del nucleo (kernel).

## Como se usan?

Se inicializan al añadir la palabra "go" antes de una funcion.