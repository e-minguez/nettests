# nettests

Simple go program and docker image to run speedtest-cli and store in a sqlite database

It runs [speedtest-cli](https://github.com/sivel/speedtest-cli) and stores the output in a sqlite database that will be used to create fancy graphs

## Build go program

```
go build -o ./nettests nettests.go
```

## Build docker image

```
docker build .
```

## Run

```
docker run -v /your/data/path/:/data:Z eminguez/nettests
```

## References
* [Most of the sql code used has been copy & pasted from siongui.github.io](https://siongui.github.io/2016/01/09/go-sqlite-example-basic-usage/)
* [Most of the run binary code used has been copy & pasted from nathan leclaire](https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/)
