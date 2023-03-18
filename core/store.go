package core

import "log"

type Provider map[string]*State

var Store = make(Provider)

func AddState(name string, initial any) *State {

	if Store[name] == nil {
		Store[name] = NewState(initial)
	} else {
		log.Printf("Error: %s existed !!", name)
	}
	return Store[name]
}
