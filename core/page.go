package core

var p *page

type page struct {
	children    []Element
	motorRender Motor
	lang        string
}

// constructor struct page
func Page(children ...Element) *page {
	defaultConf(&children)
	if p == nil {
		p = &page{
			lang:     extractLang(&children),
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
func Lang(lang string) *Ele {
	return &Ele{tag: "lang", subtype: lang}
}

/* assets */
func extractLang(c *[]Element) (res string) {
	var newchildrens []Element
	for _, item := range *c {
		if item.Tag() == "lang" {
			res = item.GetSubType()
		} else {
			newchildrens = append(newchildrens, item)
		}
	}
	c = &newchildrens
	return "es"
}

// add default configuration page
func defaultConf(children *[]Element) {
	*children = append(*children, Header(Meta(Args{Name: "theme-color", Content: "dark"})))
	*children = append(*children, Script(Args{Src: "./index.js"}))
}
