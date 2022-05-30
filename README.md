# Golang with LocalStack

## Installation and execution
This application will run with LocalStack by default.
If you want to disable it, comment lines 14-17 in main.go

start docker compose by running:
```shell
docker network create localstacknetwork
docker-compose up -d
```

then run the application:
```shell
go mod tidy
go run main.go
```

The bootstrap.sh file creates all the resource we want on localstack.
The interaction with LocalStack is seamless, only need to override the endpoint-url.
for example, when using AWSCLI:
```shell
aws --endpoint-url=http://${LOCALSTACK_HOST}:4566 s3 ls s3://
```