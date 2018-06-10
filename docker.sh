#!/bin/bash

DEMO_PATH=`pwd`
GOLANG_DOCKER_VERSION=golang:1.10
DOCKER_CONTAINER=golang-kata
APP_NAME=go-kata-practice
DOCKER_RUN_PATH=/go/src/$APP_NAME

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

  docker run -itd \
    -v $DEMO_PATH:$DOCKER_RUN_PATH \
    --name $DOCKER_CONTAINER \
    $GOLANG_DOCKER_VERSION &> /dev/null

  # docker run -itd \
  #   -v $DEMO_PATH:/go/src/go-kata-practice \
  #   --name golang-kata \
  #   golang:1.10 &> /dev/null
}

startDocker () {
  docker start $DOCKER_CONTAINER
}

enterDocker () {
  docker exec -it $DOCKER_CONTAINER /bin/bash
}

execTestKata () {
  docker exec -it $DOCKER_CONTAINER bash -c "cd $DOCKER_RUN_PATH;make coverhtml"
}

checkDockerContainerStatus () {
  check_result=`docker ps -a | grep $DOCKER_CONTAINER`

  if [ ! "$check_result" ]; then
    return 0
  fi

  check_exit=`echo $check_result | grep Exited`
  if [ ! "$check_exit" ]; then
    return 2
  fi
  return 1
}

checkDockerInstalled () {
  check_result=`which docker`

  if [ ! "$check_result" ]; then
    echo "没有检测到docker，请先安装docker之后再试"
    exit 1
  fi
}

if [ $# -ne 1 ] ; then
    usage
fi

checkDockerInstalled

checkDockerContainerStatus
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
elif [ "$1" = "test" ]; then
  if [ "$check_status" = "0" ]; then
    echo "container not exist"
    installDocker
  elif [ "$check_status" != "2" ]; then
    startDocker
  fi

  execTestKata
fi

