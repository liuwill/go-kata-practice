#!/bin/bash

usage()
{
  echo "USAGE: $CALLER [-h] COMMANDS"
  echo "       $CALLER [ --help ]"
  echo "COMMANDS:"
  echo "    install       创建docker容器"
  echo "    start         启动docker容器"
  echo "    test          运行单元测试"
  echo "    enter         进入docker容器"
  echo "Run '$CALLER COMMAND --help' for more information on a command."
  exit 1
}

installDocker () {
  echo "start install docker container"
  DEMO_PATH=`pwd`

  docker run -itd \
    -v $DEMO_PATH:/go/src/go-kata-practice \
    --name golang-kata \
    golang:1.10 &> /dev/null
}

startDocker () {
   docker start golang-kata
}

enterDocker () {
   docker exec -it golang-kata /bin/bash
}

if [ $# -ne 1 ] ; then
    usage
fi

if [ "$1" = "enter" ]; then
    enterDocker
elif [ "$1" = "install" ]; then
    installDocker
elif [ "$1" = "start" ]; then
    startDocker
fi
