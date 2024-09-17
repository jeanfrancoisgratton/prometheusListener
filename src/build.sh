#!/usr/bin/env sh

BRANCH=`git rev-parse --abbrev-ref HEAD`
BRANCH=$(echo "$BRANCH" | tr '/' '_')
BINARY=prometheusSDlistener
OUTPUT=/opt/sbin

if [ "$#" -gt 0 ]; then
    OUTPUT=$1
fi

if [ "$BRANCH" = "master" ] || [ "$BRANCH" = "main" ] || [ "$BRANCH" = "develop" ]; then
    FULLNAME="$BINARY"
else
    FULLNAME="$BINARY-$BRANCH"
fi

go build -o ${OUTPUT}/${FULLNAME} .
