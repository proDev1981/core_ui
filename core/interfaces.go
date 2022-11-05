package core

// interface motor render
type Motor interface{
  RenderElement(Element)string
  RenderPage(*page)string
}
// interface Element
type Element interface{
  render()string
  UpDate()
  childs(...Element)
  Children()[]Element
  State(s *State)Element
  SetTag(string)
  Tag()string
  SetArgs(Args)
  Args()Args
  GetSubType()string
  AddEventListener( string, func())
  SetMotorRender(Motor)
}
