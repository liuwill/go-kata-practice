#!/bin/bash

DEMO_PATH=`pwd`

docker run -itd \
  -v $DEMO_PATH:/go/src/go-kata-practice \
  --name golang-kata \
  golang:1.10 &> /dev/null
