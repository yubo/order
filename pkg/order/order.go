package order

import (
	"fmt"
	"log"
	"time"
)

const (
	DEBUG = false
)

type PickUpMsg struct {
	Order *Order
	Time  time.Duration
}

type Order struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Temp        string  `json:"temp"`
	ShelfLife   int     `json:"shelfLife"`
	DecayRate   float64 `json:"decayRate"`
	CreatedAt   int64   `json:"-"` // second
	OverflowAge int64   `json:"-"` // second
	ShelfID     int     `json:"-"`
}

func (o *Order) String() string {
	if o == nil {
		return "nil"
	}
	return fmt.Sprintf("name %s, temp %s, value %f, age %ds", o.Name, o.Temp, o.Value(), time.Now().Unix()-o.CreatedAt)
}

func (o *Order) TempID() (int, error) {
	t, ok := TempValue[o.Temp]
	if !ok || t >= TempOverflow {
		return 0, fmt.Errorf("invalid temp type %s", o.Temp)
	}
	return t, nil
}

func (o *Order) Value() float64 {
	orderAge := time.Now().Unix() - o.CreatedAt + o.OverflowAge
	if v := 1 - (o.DecayRate*float64(orderAge))/float64(o.ShelfLife); v > 0 {
		return v
	}
	return 0
}

// realAge = overflowaAge + ingle-temperature-age
// when value == 0
// => shelfLife == decayrate * orderAge * shelfDecayModifier
// => shelfLife == decayRate * (OverflowAge * 2 + ingle-temperature-age * 1)
// => shelfLife == decayRate * (OverflowAge + realAge)
// => realAge == shelfLife/decayRate - OverflowAge
// return after time
func (o *Order) CleanupTime() time.Duration {
	var realAge float64
	if o.ShelfID == TempOverflow {
		realAge = float64(o.ShelfLife) / (o.DecayRate * 2)
	} else {
		realAge = float64(o.ShelfLife)/(o.DecayRate) - float64(o.OverflowAge)
	}

	if t := (o.CreatedAt + int64(realAge)) - time.Now().Unix(); t > 0 {
		return time.Duration(t) * time.Second
	}
	return 0
}

func info(format string, a ...interface{}) {
	log.Output(2, fmt.Sprintf(format, a...))

}

func debug(format string, a ...interface{}) {
	if DEBUG {
		log.Output(2, fmt.Sprintf(format, a...))
	}
}
