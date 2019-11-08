# gatic

[![Docker pulls](https://img.shields.io/docker/pulls/twinproduction/gatic.svg)](https://cloud.docker.com/repository/docker/twinproduction/gatic)

Very small static web server in Go.


## Usage

All you have to do is create a Dockerfile that uses this image, and copy the static files you want to expose
in the `static` folder of the Docker image:

```dockerfile
FROM twinproduction/gatic
ADD static ./static
```

See the [example](example) folder for an example.


## Docker

Building the Docker image is done as following:

```
docker build . -t gatic
```

You can then run the container with the following command:

```
docker run -p 8080:8080 --name gatic gatic
```


## Running the tests

```
go test ./... -mod vendor -cover
```
