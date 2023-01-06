package core

type Provider struct {
	dataBase map[string]*State
}

// constructor
func newProvider() *Provider {
	return &Provider{dataBase: make(map[string]*State)}
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
func (p *Provider) AddState(s *State) {
	p.dataBase[s.name] = s
}
