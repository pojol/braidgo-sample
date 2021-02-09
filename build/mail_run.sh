#! bin/bash

docker run -d  -p 14301:14301/tcp braid-sample/mail:latest \
    -consul http://172.17.0.1:8500 \
