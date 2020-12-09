# Docker

* [x] [Installation in Ubuntu](https://docs.docker.com/engine/install/ubuntu/)
  * [x] [Got permission denied issue](https://stackoverflow.com/a/48957722/256002)

```bash
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassowrd -d postgres:9.6
```

* Container vs Image
  * Container is a running environment for an IMAGE
  * Container provides file system, environment configs for an application image (postgres, redis, mongo)
  * Container has a port binded to it in order to talk with the application that runs within it

Commands

```bash
docker pull IMAGE_NAME # downloads the image from the Docker Hub
docker images # lists all the images on the local machine
docker run IMAGE_NAME # pulls the image and star a container running this image in an ATTACHED mode (to terminate press Ctrl+C)
docker run -p6001:6379 IMAGE_NAME # runs a container with the port of the host bound to the port of the docker container (HOST_PORT:CONTAINER_PORT)
docker ps # list running containers
docker ps -a # list running or NOT RUNNING containers 
docker stop CONTAINER_ID [CONTAINER_ID2, ....] # to stop the container
docker start CONTAINER_ID [CONTAINER_ID2, ....] # to start the container
docker logs CONTAINER_ID | CONTAINER_NAME # shows the logs of the container
docker exec -it CONTAINER_ID | CONTAINER_NAME /bin/bash # attaches to the bash inside of the container
docker system prune #clean up any resources — images, containers, volumes, and networks — that are dangling (not associated with a container)
docker network ls # list of networks
docker network create NEW_NETWORK_NAME
```

```bash
# EXAMPLES
docker run -d -p6000:6379 --name redis-older redis:4.0 # runs a container from the image binding host OS port 6000 to docker container post 6379, gives this container name redis:4.0
docker exec -it 79899b6a6c08 /bin/bash # executes interactive (-it) terminal (/bin/bash) inside of the container 79899b6a6c08
env # executed inside of the -it prints environment variables
exit # to exit terminal
```

```bash
# creating a sample mongodb-backed application
docker run -p 27017:27017 -d -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password --name mongodb --network mongo-network mongo
docker run -d \
> -p 8081:8081 \
> -e ME_CONFIG_MONGODB_ADMINUSERNAME=admin \
> -e ME_CONFIG_MONGODB_ADMINPASSWORD=password \
> --net mongo-network \
> --name mongo-express \
> -e ME_CONFIG_MONGODB_SERVER=mongodb \
> mongo-express
```
