DOCKER_DEFAULT_PLATFORM=linux/amd64 docker build  -t harbor.wns8.io/public/ttv-base:1.0.0  -f  ./.deploy/ttv_Dockerfile .

# DOCKER_DEFAULT_PLATFORM=linux/amd64  docker run -itd  harbor.wns8.io/public/ttv-base:1.0.0  /bin/bash
