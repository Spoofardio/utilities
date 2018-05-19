# Docker Redo

Small script to recreate a docker container

## How to use it

### Get Help
```
source docker-redo.sh -h
```
### General Form
```
source docker-redo.sh -n <NAME> -r "<ARGS>" -i <IMAGE>
```
Required Arguments:
```
IMAGE(-i, --image): the docker image. this is required and there is no default value.
```
Optional Arguments:
```
NAME(-n, --name: the name of the container. default value is `redoContainer`
ARGS: the docker run arguments. default value is `-d`
```

### Example
```
source docker-redo.sh -n rabbitmq -i rabbitmq:3.7 -r "-d --hostname rabbitmq -p 5672:5672 -p 15672:15672 -v rabbitmq-home:/var/lib/rabbitmq -v rabbitmq-config:/etc/rabbitmq/rabbitmq.config --restart unless-stopped" 
```

## Different Modes

currently there is only one mode which is the default.
```
stop the container
remove the container
delete the image
repull the image
recreate the container 
```

## Authors

* **Zachary Spofford** - *Initial work* - [Spoofardio](https://github.com/Spoofardio)