package order

import (
	"fmt"
	"sync"
	"time"
)

const (
	TempHot = iota
	TempCold
	TempFrozen
	TempOverflow
	TempSize
)

var (
	TempName = map[int]string{
		0: "hot",
		1: "cold",
		2: "frozen",
		3: "overflow",
	}

	TempValue = map[string]int{
		"hot":      0,
		"cold":     1,
		"frozen":   2,
		"overflow": 3,
	}
)

// shelves recv order and palce the order on the best-available shelf
type Kitchen struct {
	sync.RWMutex
	shelves    [TempSize]map[string]*Order
	shelvesCap [TempSize]int
	orderCh    chan *Order // receive
	pickUpCh   chan *Order // pickup
}

func NewKitchen(orderCh, pickUpCh chan *Order, caps [TempSize]int) *Kitchen {
	ret := &Kitchen{
		shelves:    [TempSize]map[string]*Order{},
		shelvesCap: caps,
		orderCh:    orderCh,
		pickUpCh:   pickUpCh,
	}
	for k, _ := range ret.shelves {
		ret.shelves[k] = make(map[string]*Order)
	}
	return ret
}

func (p *Kitchen) PickUp(o *Order) error {
	p.Lock()
	defer p.Unlock()

	i, err := o.TempID()
	if err != nil {
		return err
	}

	if _, ok := p.shelves[i][o.Id]; ok {
		delete(p.shelves[i], o.Id)
		debug("delete shelves[%s][%s] %d/%d",
			TempName[i], o.Id, len(p.shelves[i]),
			p.shelvesCap[i])
		return nil
	}

	if _, ok := p.shelves[TempOverflow][o.Id]; ok {
		delete(p.shelves[TempOverflow], o.Id)
		debug("delete shelves[%s][%s] %d/%d",
			TempName[i], o.Id, len(p.shelves[i]),
			p.shelvesCap[i])
		return nil
	}

	return fmt.Errorf("not found order %s", o)
}

func (p *Kitchen) Place(o *Order) error {
	p.Lock()
	defer p.Unlock()

	i, err := o.TempID()
	if err != nil {
		return err
	}

	defer func() {
		debug("pickUpCh <- %s", o)
		p.pickUpCh <- o
	}()

	if n := len(p.shelves[i]); n < p.shelvesCap[i] {
		p.shelves[i][o.Id] = o
		o.ShelfID = i
		info("place shelves[%s] %d/%d %s",
			TempName[i], len(p.shelves[i]), p.shelvesCap[i], o)
		return nil
	}

	o.ShelfID = TempOverflow
	if n := len(p.shelves[TempOverflow]); n < p.shelvesCap[TempOverflow] {
		p.shelves[TempOverflow][o.Id] = o
		debug("insert shelves[%s] %d/%d %s shelves[%s] is full ",
			TempName[TempOverflow], len(p.shelves[TempOverflow]),
			p.shelvesCap[TempOverflow], o, TempName[i])
		return nil
	}

	if ok := p.cleanUpOveflow(); ok {
		p.shelves[TempOverflow][o.Id] = o
		debug("insert shelves[%s] %d/%d %s shelves[%s] is full ",
			TempName[TempOverflow], len(p.shelves[TempOverflow]),
			p.shelvesCap[TempOverflow], o, TempName[i])
		return nil
	}

	// map is random, so, drop the first one
	for k, v := range p.shelves[TempOverflow] {
		delete(p.shelves[TempOverflow], k)
		info("discarded order %s overflow shelf is full", v)
		debug("delete shelves[%s][%s] %d/%d",
			TempName[TempOverflow], k, len(p.shelves[TempOverflow]),
			p.shelvesCap[TempOverflow])
		break
	}

	p.shelves[TempOverflow][o.Id] = o
	return nil
}

// the overflow shelf is full,
// an existing order of your choosing on the overflow should be moved
// to an allowable shelf with room
func (p *Kitchen) cleanUpOveflow() bool {
	for i := 0; i < TempOverflow; i++ {
		if len(p.shelves[i]) < p.shelvesCap[i] {
			// find the temp order from overflow, and move to this shelf
			for k, v := range p.shelves[TempOverflow] {
				if TempValue[v.Temp] == i {
					delete(p.shelves[TempOverflow], k)

					// move from overflow -> i
					p.shelves[i][v.Id] = v
					v.OverflowAge = time.Now().Unix() - v.CreatedAt
					v.ShelfID = i
					info("move order %s from overflow shelf to shelves[%s]",
						v, TempName[i])
					return true
				}
			}
		}
	}
	return false
}

func (p *Kitchen) String() string {
	p.RLock()
	defer p.RUnlock()

	return fmt.Sprintf("%s %d %s %d %s %d %s %d",
		TempName[TempHot], len(p.shelves[TempHot]),
		TempName[TempCold], len(p.shelves[TempCold]),
		TempName[TempFrozen], len(p.shelves[TempFrozen]),
		TempName[TempOverflow], len(p.shelves[TempOverflow]),
	)
}

func (p *Kitchen) Start() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case o := <-p.orderCh:
			info("kitchen recv order %s", o)
			p.Place(o)
		case <-ticker.C:
			debug("shelves %s", p)
		}
	}

}
