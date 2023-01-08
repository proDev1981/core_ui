package core

var p *page

type page struct {
	children    []Element
	motorRender Motor
	lang        string
}

// constructor struct page
func Page(lang string, children ...Element) *page {
	if p == nil {
		p = &page{
			lang:     lang,
			children: children,
		}
	}
	return p
}

// render page
func (p *page) Render() (res string) {
	return p.motorRender.RenderPage(p)
}

// getter children page
func (p *page) Children() []Element {
	return p.children
}

// setter motor render
func (p *page) SetMotorRender(m Motor) {
	p.motorRender = m
}

// getter events listeners
func (p *page) GetMotorRender() Motor {
	return p.motorRender
}

func Link(args Args) *Ele {
	args.Rel = "stylesheet"
	return &Ele{tag: "link", args: args}
}
func Meta(args Args) *Ele {
	if args.Name == "theme-color" {
		args.id = "color-title"
	}
	return &Ele{tag: "meta", args: args}
}
func PageTitle(args Args) *Ele {
	return &Ele{tag: "title", args: args}
}
func Header(children ...Element) *Ele {
	return &Ele{tag: "head", children: children}
}
func Script(args Args) *Ele {
	return &Ele{tag: "script", args: args}
}
