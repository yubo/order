package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/yubo/order/pkg/order"
	"github.com/yubo/order/pkg/order/emulator"
)

var (
	rate     int
	filename string
)

func main() {
	flag.StringVar(&filename, "filename", "./var/orders.json", "orders.json path")
	flag.IntVar(&rate, "rate", 2, "the rate of the orders receive")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.Lshortfile)

	orderCh := make(chan *order.Order, 10)
	pickUpCh := make(chan *order.Order, 10)

	kitchen := order.NewKitchen(orderCh, pickUpCh, [order.TempSize]int{10, 10, 10, 15})

	emu, err := emulator.NewEmulator(rate, filename, kitchen, orderCh, pickUpCh)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	go kitchen.Start()
	go emu.Start()
	select {}
}
