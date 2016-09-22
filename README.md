# nettests

Simple go program and docker image to run speedtest-cli and store in a sqlite database

## Build go program
```
go build -o ./nettests nettests.go
docker build .
```

## Run

```
docker run -v /your/data/path/:/data:Z eminguez/nettests
```
