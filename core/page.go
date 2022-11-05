package core

type page struct{
  children []Element
  motorRender Motor
}
// constructor struct page
func Page(t Motor,children ...Element)*page{
  p:= &page{ children,t }
  for _,item:= range p.children{
    item.SetMotorRender(t)
  }
  p.motorRender.SetPage(p)
  return p
}
// render page
func (p *page) Render()(res string){
  return p.motorRender.RenderPage()
}
// getter children page
func (p *page) Children()[]Element{
  return p.children
}

// falta por implementar
func Styles(path string)*Ele{
  return &Ele{tag:"style"}
}
func Header(children ...*Ele)*Ele{
  return &Ele{tag:"header"}
}
