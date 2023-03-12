package core

type Args struct {
	Key         string
	Class       string
	Name        string
	id          string
	Type        string
	Value       string
	Src         string
	Alt         string
	List        string
	Each        string
	FallBack    Element
	State       *State
	Store       Provider
	Events      Listener
	Direction   string
	Rel         string
	Href        string
	Link        *string
	Width       string
	Height      string
	Max         int
	Min         int
	Charset     string
	Content     string
	Placeholder string
}
