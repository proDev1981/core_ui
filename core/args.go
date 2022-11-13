package core

type Args struct {
	Class     string
	Name      string
	id        string
	Style     Style
	Type      string
	Value     string
	Src       string
	Alt       string
	State     *State
	Events    Listener
	Direction string
	Rel       string
	Href      string
}
