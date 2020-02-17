

default: kvserver kvclient

clean:
	rm -f kvserver kvclient mem.log

kvserver: server/*.go
	go build -o kvserver server/*.go

kvclient: client/*.go
	go build -o kvclient client/*.go

test: default
	scripts/simple-test.sh
