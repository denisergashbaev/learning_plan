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
```
