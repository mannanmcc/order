### Run the application in docker container

$ docker build -t mannanmcc/order:v1 .
$ docker run -p 50052:50051 --net docker-network -e CONFIG_FILE=build/config.yaml mannanmcc/order:v1

Make sure we created a docker network first using docker network create command and launce stock api using similar command above before starting above command as this API has dependency.