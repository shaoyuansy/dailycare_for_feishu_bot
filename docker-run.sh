#!/bin/bash

docker stop informal_bot > /dev/null 2>&1
docker rm informal_bot > /dev/null 2>&1
image_tag=(`cat version | grep "version" -m 1 | cut -d= -f2 | sed 's/\"//g'`)
docker run -d -p 8199:8199 --name="informal_bot" -v /var/log/informal:/var/log/informal -v /app/dailycare_for_feishu_bot/config:/var/www/informal/config informal_bot:"v${image_tag}"
docker logs -f informal_bot