#!/bin/bash
gf build

image_tag=(`cat version | grep "version" -m 1 | cut -d= -f2 | sed 's/\"//g'`)

echo "Building InformalBot docker image tagged as [v$image_tag]"

docker build --tag="informal_bot:v$image_tag" . \
  && echo "Built informal_bot image successfully tagged as informal_bot:v$image_tag" \
  && docker images "informal_bot:v$image_tag"
