# webapp

Simple Go webapp which sends data to InfluxDB.

## usage

```bash
$ export APP_PORT=3001
$ export INFLUX_ADDR=http://localhost:8086
$ export DATABASE=webapp
$ go run main.go
```

## docker

```bash
$ cd docker
$ ./build.sh
$ docker-compose up -d
```