#! bin/bash

docker run -d -p 14201:14201/tcp braid-sample/base:latest \
    -consul http://172.17.0.1:8500 \
    -nsqlookup 172.17.0.1:4161 \
    -nsqd 172.17.0.1:4150 \
