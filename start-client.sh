#!/bin/sh
GLIDE=`which glide`

if [[ $GLIDE == "" ]]; then
    echo "\`glide\` is not installed, please install it first!"
    exit 0
fi

cd src

$GLIDE build-client

../bin/client
