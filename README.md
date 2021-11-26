# PAPER
This project is example go.

## Development

### Run App
```
go run .
```

### Run Test
```
go test -v ./...
```

## Dockerize

### Build
```
docker build -t paper .
```

### Run
```
docker run --rm -p 8080:8080 paper
```

### DockerHub Pull

#### Link:
![DockerHub Link](https://hub.docker.com/r/openfile/paper)

```
docker pull openfile/paper:master
```
