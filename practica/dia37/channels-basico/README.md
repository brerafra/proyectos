# Elementos básicos de channels

1. Crear un canal: se utiliza la funcion make
    -   cana:=make(chan string)

2. Enviar datos: Se utiliza el operador <- apuntando hacia el canal
    -   canal <- "Dato a enviar"

3. Recibir datos: Se utiliza el operador <-  apuntando hacia el canal.
    - variable := <-canal