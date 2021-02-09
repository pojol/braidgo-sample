#! bin/bash

docker run -d  -p 14101:14101/tcp braid-sample/login:latest \
    -consul http://172.17.0.1:8500 \
    -nsqlookup 172.17.0.1:4161 \
    -nsqd 172.17.0.1:4150 \
