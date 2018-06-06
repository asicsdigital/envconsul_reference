# Build

```sh
$ mkdir -p ~/go/github.com/asicsdigital
$ pushd ~/go/github.com/asicsdigital
$ git clone https://github.com/asicsdigital/envconsul_reference
$ cd envconsul_reference
$ go build
```

# Run
```sh
$ HELLO_NAME=Steve ./envconsul_reference &
$ curl http://localhost:8080
```

Pass env var `PORT` to listen on a different port.
