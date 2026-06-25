# DOCKER

Docker es una herramienta de contenerización que te permite desplegar y configurar servicios en entornos aislados, para gestionar diversos aspectos de esta plataforma y de los servicios sobre ella, utilizas principalmente varios comandos docker.

# Comandos docker

## Comandos de construcción:

Crea una magen apartir de un archivo docker

- docker build          Construye una imagen a partir de un docker file en el directorio actual
- docker build "github" Construye una imagen desde un repositorio GIT remoto.
- docker build -t
  imagenname/tag        Construye y etiqueta una imagen para facilitar su seguimiento.
- docker build https
://yourserver/file.tar.gz   Crea una imagen apartir de un archivo tar remoto
- docker build -t imagen:
1.0-<<EOFFROM bussyboxRUN
echo "hola mundo" EOF   Construye una imagen mediante un archivo dockerfile que se pasa a travéz de STDIN

## Comandos de limpieza:

Elimina imágenes y volúmenes no utilizados para liberar espacio

- docker image prune        borra una imagen no utilizada
- docker image prune -a     borra todas las imágenes que no esten siendo utilizadas por contenedores
- docker system prune       Elimina todos los contenedores detenidos, todas las redes no utilizadas por los
                            contendores, todas las imagenes colgadas y toda la cache de construcción.
- docker image rm image     elimina una imagen
- docker rm container       Elimina un contenedor en ejecución
- docker kill $(docker
  ps -q)                    Detiene todos los contenedores en ejecución
- docker swarm leave        Deja un enjambre
- docker stack rm stackname elimina un enjambre
- docker volume rm $( 
  docker volume ls -f
  dangling=true -q)         Elimina todos los volúmenes colgados
- docker rm $(docker ps
  -a -q)                    Elimina todos los contenedores parados

## Comandos de interacción con contenedores:

Gestiona los contenedores y comunícate con ellos

- docker start container    Inicia un nuevo contenedor
- docker stop container     Detiene un contenedor
- docker pause container    Pausa un contenedor
- docker unpause container  despausa
- docker restart container  reinicia un contenedor
- docker wait container     bloquea un contenedor
- docker export container   Exporta el contenido del contendeor a un archivo tar.
- docker attach container   Se une a un contenedor en ejecución
- docker commit -m "commit
  message" -a "author" container
  username/image_name: tag  Guarda un contenedor en ejecución como una imagen.
- docker logs -ft container Sigue los registros de contenedores
- docker exec -ti container
  script.sh                 Ejecuta un comando en un contenedor
- docker commit container
  image                     Crea una nueva imagen a partir de un contenedor
- docker create image       Crea un nuevo contenedor a partir de una imagen.

## Inspección de contenedores

analiza y comprueba detalles de los contenedores

- docker ps         Lista todos los contenedores en ejecución
- docker ps -a      Lista todos contenedores
- docker diff 
    container       Inspecciona los cambios en los directorios y archivos del sistema de archivos del contenedor
- docker top
    container       Muestra todos los precesos en ejecución de un contenedor existente
- docker inspect    
    container       Muestra información de bajo nivel sobre un conenedor
- docker logs 
    container       Reúne los registros de un contenedor
- docker stats 
    container       Muestra las estadísticas de uso de los recursos de un contenedor

## Comandos de gestión de imagenes:

administra imágenes

- docker image ls       Lista imagenes
- docker image rm mysql Elimina una imagen
- docker tag image tag  etiqueta una imagen
- docker history image  muestra el historial de imágenes
- docker inspect image  muestra información de bajo nivel sobre una imagen


## ejecutar comandos: 

construye un contenedor a partir de una imagen y cambia su configuración

El comando run en docker se utiliza para crear contenedores apartir de imagenes proporcionadas, la sintaxis es:

- docker run (options) image (command) (arg...)

se pueden usar las siguientes banderas para modificar el comportamiento del comando

-detach, -d     ejecuta un contenedor en segundo plano e imprime el ID del contenedor
-env, -e        establece variables de entorno
-hostname, -h   estable un nombre del host a un contenedor
-label, -l      Crea una etiqueta de metadatos para un contenedor
-name           asgina un nombre a un contenedor
-network        Conecta un contenedor a una red
-rm             retira un contenedor cuando se detenga
-read-only      Establece el sistema de archivos del contenedor como solo lectura
-workdir, -w    establece un directorio de trabajo en un contenedor


## Comandos de registro:

interactua con un registrod e imágenes Docker remoto, como Docker hub.

- docker login      Accede a un registro
- docker logut      Sal de un registro
- docker pull mysql Extrae una imagen de un registro
- docker push repo/
rhel-httpd:latest   Envia una imagen a un registro
- docker search term    Busca en docker hub imagenes con el termino especificado.

## Comandos de servicio:

gestiona todos los aspectos de los servicios Docker

- docker service ls         Lista todos los servicios que se ejecutan en un enjambre
- docker stack services
    stackname               Lista todos los servicios en ejecución
- docker service ps
    servicename             Lista las tareas de un sercicio
- docker service update
    servicename             Actualiza un servicio
- docker service create
    image                   Crea un nuevo servicio
- docker service scale
    servicename=10          Escala uno o más servicios replicados
- docker service logs
    stackname servicename   lita todos los registros de servicio

## comandos de red: 

configura, gestiona e interactua con la red docker

- docker network create
    networkname             Crea una nueva red
- docker network rm
    networkname             Elimina una red epecificada
- docker network ls         lista todas las redes
- docker network connect
    networkname container   conecta un contenedor a una red
- docker network disconnect
    networkname container   desconecta un contenedor a una red
- docker network inspect
    networkname             muestra información detallada sobre una red