.PHONY: clean run

all: order-emulator

order-emulator: $(shell find . -name "*.go")
	GO111MODULE=on go build -o $@ ./cmd/order-emulator

clean:
	rm -f order-emulator

run: order-emulator
	./order-emulator --filename ./var/orders.json --rate 2
