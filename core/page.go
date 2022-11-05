package core

type page struct{
  children []Element
  motorRender Motor
}
// constructor struct page
func Page(children ...Element)*page{
  p:= &page{ children:children }
  return p
}
// render page
func (p *page) Render()(res string){
  return p.motorRender.RenderPage(p)
}
// getter children page
func (p *page) Children()[]Element{
  return p.children
}
// setter motor render
func (p *page) SetMotorRender(m Motor){
  p.motorRender= m
}

// falta por implementar
func Styles(path string)*Ele{
  return &Ele{tag:"style"}
}
func Header(children ...*Ele)*Ele{
  return &Ele{tag:"header"}
}

