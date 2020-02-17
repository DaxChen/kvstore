# kvstore

CS739 Project 1: fault-tolerant key value store using gRPC

# installing Go and dependencies

If you already have go 13.8 and go tools (eg. go mod) installed, you can skip this.

Follow the offical guild to install `go` and go-tools (eg. `go mod`).
If on `ubuntu 16.04` to `18.04`, you can take a look at `scripts/install-go-ubuntu.sh`, which
downloads and installs go from official website and setup PATH and GOPATH for you.

The module dependencies are managed by go modules.

Make sure you have go modules installed and go should automatically download the needed dependencies when you run `go run` or `go build`.

(You should clone this repo outside of GOPATH, because by default go modules are disabled inside GOPATH. You can either move this folder outside of GOPATH or enable go modules everywhere.)

# Compiling and `make` commands

- `make`: compile both server and client binary executables using `go build`.
- `make clean`: to remove `kvserver`, `kvclient`, and the database file `mem.log`.
- `make test`: will run `scripts/simple-test.sh`, will start server and clients on same machine, load some data, run some basic operations for about 20 seconds, and kill server.

# run test scripts

Under the `scripts` directory, there are several scripts to run our experiments.

for example run `scripts/exp1-server.sh` to start server, and `scripts/exp1-client.sh` to start client for to run experiment 1, it will ask you to provide some numbers to use.
