package core

import (
	"sync"
	"time"

	"github.com/castillobg/pong/brokers"
)

type pongs struct {
	sync.Mutex
	count int
}

var p *pongs

func Listen(broker brokers.BrokerAdapter, pings chan []byte, delay int) {
	p = &pongs{}
	go func() {
		// Listens for ping events
		for range pings {
			// If a ping arrives, wait for 2 sec. then respond with a pong.

			go func() {
				time.Sleep(time.Duration(delay) * time.Second)
				broker.Publish("pong", "pongs")
				p.Lock()
				p.count++
				p.Unlock()
			}()
		}
	}()
}

func Pongs() int {
	return p.count
}
