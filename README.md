# recruitment-tool-web
a simple web app for formatting recruitment urls because Discord will no longer
do it for us

## build & run
1. `go build -o recruitment-tool-web`
2. `./recruitment-tool-web`

## docker build & run
1. `docker build -t recruitment-tool-web .`
2. `docker run -p 8080:8080 --rm recruitment-tool-web`
