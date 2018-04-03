#!/bin/bash
usage="use like this: \$source docker-redo.sh -n NAME -r \"ARG ARG\" -i IMAGE"
arg_error="ERROR: invalid arguments. for help run \$source docker-redo.sh -h"
image_error="ERROR: you must specify an image"

name="redoContainer"
runArgs="-d"
image=""

while [ "$1" != "" ]; do
  case $1 in
    -r | --runargs )  shift
                      runArgs=$1
                      ;;
    -n | --name )     shift
                      name=$1
                      ;;
    -i | --image)     shift
                      image=$1
                      ;;
    -h | --help )     echo "$usage"
                      return
                      ;;
    * )               echo "$arg_error"
                      return
  esac
  shift
done

if ["$image" == ""] 
then
  echo "$image_error"
  return
fi

docker stop $name
docker rm $name
docker rmi $image
docker pull $image
docker run $runArgs --name $name $image
