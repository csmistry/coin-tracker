# coin-tracker


## Build backend

Make sure you are in the root directory

1. build backend image
```
docker build -t go-web-app .
```
2. Run the docker container on port `:8080`
```
docker run -p 8080:8080 go-web-app
```