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

checkDockerStatus () {
  check_result=`docker ps -a | grep golang-kata`

  if [ ! "$check_result" ]; then
    return 0
  fi
  return 1
}

if [ $# -ne 1 ] ; then
    usage
fi

checkDockerStatus
check_status=$?

if [ "$1" = "install" ]; then
  if [ "$check_status" != "0" ]; then
    echo "容器已经存在"
    exit 0
  fi
  installDocker
  exit 0
fi


if [ "$1" = "start" ]; then
  if [ "$check_status" = "2" ]; then
    echo "container is working"
    exit 0
  elif [ "$check_status" = "0" ]; then
    echo "container not exist"
    installDocker
  fi
  startDocker
elif [ "$1" = "enter" ]; then
  if [ "$check_status" = "0" ]; then
    echo "container not exist"
    installDocker
  elif [ "$check_status" != "2" ]; then
    startDocker
  fi

  enterDocker
fi
