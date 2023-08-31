# Go Fetch Weather
Go application that fetches weather alerts and exposes data via REST API 

## 🙇 Application Requirement
1. Install Docker Desktop via https://docs.docker.com/desktop/
2. Install Go via https://go.dev/doc/install

Once installed, Please check via below commands
```zsh   
go version
```

```zsh   
docker version
```

## 🛠️ Start the application locally

1. Clone the repository
2. Run `go get .` to install all the required dependencies
3. Open a terminal and run comand below

```zsh
make run
```

## How run inside a Docker container,

Open a terminal and run comand in order to build an image

```zsh
docker build -t go-fetch-weather:local .
```

Comand below will run the image inside a container

```zsh
docker run --rm -p 8081:8081 -it go-fetch-weather:local
```

## How to use

Open a terminal and run comand below

```zsh
curl --location 'localhost:8081/alerts?area=NY'
```

## Resources
- https://medium.com/@faruqisan/how-to-create-an-api-client-in-golang-a5d3f56b4080
- https://stackoverflow.com/questions/43976140/check-errors-when-calling-http-responsewriter-write
- https://askgolang.com/cannot-convert-data-type-interface-to-type-string-need-type-assertion/
- https://www.golangprograms.com/golang-http-get-request-with-parameters.html
- https://pkg.go.dev/github.com/paulmach/orb/geojson#section-readme
- https://github.com/paulmach/orb/tree/master/geojson
- https://golangbyexample.com/net-http-package-get-query-params-golang/
- https://golangbyexample.com/json-response-body-http-go/
- https://earthly.dev/blog/golang-http/
- https://gist.github.com/tomnomnom/52dfa67c7a8c9643d7ce
- https://golangtutorial.dev/tips/http-post-json-go/
- https://tutorialedge.net/golang/consuming-restful-api-with-go/
