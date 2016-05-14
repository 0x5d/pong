package core

import (
	"sync"

	"github.com/castillobg/pong/brokers"
)

type pongs struct {
	sync.Mutex
	count int
}

var p *pongs

func Listen(broker brokers.BrokerAdapter, pings chan []byte) {
	p = &pongs{}
	go func() {
		// Listens for pong events
		for range pings {
			// If a pong arrives, respond with a ping.
			broker.Publish("pong", "pongs")
			p.Lock()
			p.count++
			p.Unlock()
		}
	}()
}

func Pongs() int {
	p.Lock()
	defer p.Unlock()
	return p.count
}
