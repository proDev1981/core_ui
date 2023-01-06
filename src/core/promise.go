package core

import "fmt"

type PROMISE struct {
	id       string
	state    string
	data     string
	resolved func(*PROMISE)
	reyect   func(*PROMISE)
	motor    Motor
}

// contructor promise
func NewPromise(m Motor) *PROMISE {
	p := &PROMISE{state: "procesing", motor: m}
	p.id = fmt.Sprintf("%p", p)
	return p
}

// if promise is resolved call fun pass of argument
func (p *PROMISE) Then(res func(*PROMISE)) *PROMISE {
	p.resolved = res
	p.motor.NewServer().GetSocket().AddPromise(p)
	return p
}

// if promiese is reyect call func pass of argument
func (p *PROMISE) Catch(rey func(*PROMISE)) *PROMISE {
	p.reyect = rey
	return p
}

// create conection witch chanel data promise
func (p *PROMISE) Await() chan string {
	c := make(chan string)
	p.Then(func(res *PROMISE) {
		c <- res.GetData()
	})
	return c
}

// change state of promise
func (p *PROMISE) ChangeState(value string) *PROMISE {
	p.state = value
	return p
}

// set data
func (p *PROMISE) SetData(data string) *PROMISE {
	p.data = data
	return p
}

// get data
func (p *PROMISE) GetData() string {
	return p.data
}
