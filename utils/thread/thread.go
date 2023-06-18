package thread

import (
	"github.com/emicklei/go-restful/v3/log"
)

type Handler func(any) error

type Thread struct {
	Handler func(any) error
	Val     chan any
}

func New(handler Handler) *Thread {
	p := &Thread{
		handler,
		make(chan any),
	}
	go p.run()
	return p
}
func (p *Thread) Push(val any) {
	if p == nil {
		return
	}
	go func() {
		p.Val <- val
	}()
}
func (p *Thread) run() {
	for {
		val := <-p.Val
		err := p.Handler(val)
		if err != nil {
			log.Print("error email delivery", err)
		}
	}
}
