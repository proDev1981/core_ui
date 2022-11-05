package core

type Motor interface{
  RenderElement(Element)string
  RenderPage()string
  SetPage(*page)
}
