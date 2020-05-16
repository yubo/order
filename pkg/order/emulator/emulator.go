package emulator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/yubo/order/pkg/order"
)

// emulate order and send the order to Shelves's channel
// and create a couriers to pick up the order between 2-6 seconds later
type Emulator struct {
	OrderRate int
	Orders    []*order.Order
	kitchen   *order.Kitchen
	orderCh   chan *order.Order
	pickUpCh  chan *order.Order
}

func NewEmulator(rate int, filename string, kitchen *order.Kitchen, orderCh, pickUpCh chan *order.Order) (*Emulator, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	ret := &Emulator{
		OrderRate: rate,
		Orders:    []*order.Order{},
		kitchen:   kitchen,
		orderCh:   orderCh,
		pickUpCh:  pickUpCh,
	}

	err = json.Unmarshal(buf, &ret.Orders)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (e *Emulator) Start() {
	info("emulator start, %d order per second, total %d",
		e.OrderRate, len(e.Orders))

	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(e.OrderRate))
		defer ticker.Stop()

		total := len(e.Orders)
		for k, v := range e.Orders {
			v.CreatedAt = time.Now().Unix()
			<-ticker.C
			info("Send Order %d/%d %s", k+1, total, v)
			e.orderCh <- v
		}
	}()

	for {
		select {
		case o := <-e.pickUpCh:
			debug("pickUpCh -> order %s", o)
			e.pickUp(o)
		}
	}
}

func (e *Emulator) pickUp(o *order.Order) {
	// for courier
	// random time (2~6 second)
	t := rand.Intn(4000) + 2000
	time.AfterFunc(time.Duration(t)*time.Millisecond, func() {
		if err := e.kitchen.PickUp(o); err != nil {
			debug("[err] courier pickUp err %v", err)
		} else {
			info("PickUp the order %s", o)
		}
	})

	// for cleanup
	time.AfterFunc(o.CleanupTime(), func() {
		if o.OverflowAge > 0 {
			// move from overflow shelf
			time.AfterFunc(o.CleanupTime(), func() {
				if err := e.kitchen.PickUp(o); err != nil {
					debug("cleanup err %v", err)
				} else {
					info("CleanUp order %s", o)
				}
			})
			return
		}

		if err := e.kitchen.PickUp(o); err != nil {
			debug("cleanup err %v", err)
		} else {
			info("CleanUp order %s", o)
		}
	})
}

func info(format string, a ...interface{}) {
	log.Output(2, fmt.Sprintf(format, a...))

}

func debug(format string, a ...interface{}) {
	if order.DEBUG {
		log.Output(2, fmt.Sprintf(format, a...))
	}
}
