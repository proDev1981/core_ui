package core

type Provider struct {
	dataBase map[string]*State
}

// constructor
func NewProvider() *Provider {
	return &Provider{make(map[string]*State)}
}

// getter state for map
func (p *Provider) GetState(name string) *State {
	res := p.dataBase[name]
	if res != nil {
		return res
	}
	return nil
}

// setter state for map
func (p *Provider) AddState(name string, s *State) {
	p.dataBase[name] = s
}
