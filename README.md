## ascii-art-web-dockerize

# Project

You will have a go web server for ascii-art. You can type in some text, choose the font from 3 fonts `standard`, `shadow `or `tinkertoy` and get nice output on the screen or as one of the text files from `.txt` or `.doc` for download. :) 

![ascii-art](/static/images/ascii-art.png)

## Instructions how to run docker file

### To build the Docker image

![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) **Mac**

Navigate to the directory containing the Dockerfile and run the following command:

```
docker image build -t go-web-server .
```

This will build the Docker image and give it the tag go-web-server. You can then run the image using the docker run command:

```
docker container run -p 8080:8080 -it go-web-server
```
**or use script to build the images and containers**
```
sh dockerBuild.sh
```

![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) **Linux**

```
sudo docker image build -t go-web-server .
```

This will build the Docker image and give it the tag go-web-server. You can then run the image using the docker run command:

```
sudo docker container run -p 8080:8080 -it go-web-server
```

**or use script to build the images and containers**
```
sh dockerBuildSudo.sh
```

This will start the Go web server and bind it to port 8080 on the host machine.


### To remove Docker image use:

![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) **Mac**

stop all running containers
```
echo 'Stopping running containers (if available)...'
docker stop $(docker ps -aq)
```
remove all stopped containers
```
echo 'Removing containers ..'
docker rm $(docker ps -aq)
```
remove all images
```
echo 'Removing images ...'
docker rmi $(docker images -q)
```
remove all stray volumes if any
```
echo 'Removing docker container volumes (if any)'
docker volume rm $(docker volume ls -q)
```
**or use script to remove build images and containers**
```
sh dockerClearImages.sh
```

![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) **Linux**

stop all running containers
```
echo 'Stopping running containers (if available)...'
sudo docker stop $(sudo docker ps -aq)
```
remove all stopped containers
```
echo 'Removing containers ..'
sudo docker rm $(sudo docker ps -aq)
```
remove all images
```
echo 'Removing images ...'
sudo docker rmi $(sudo docker images -q)
```
remove all stray volumes if any
```
echo 'Removing docker container volumes (if any)'
sudo docker volume rm $(sudo docker volume ls -q)
```

**or use script to remove build images and containers**
```
sh dockerClearImagesSudo.sh
```
