package core

import "fmt"
import "strings"

type Html struct{
  Page *page
}
// contructor
func NewHtml()*Html{
  return &Html{}
}
// setter page in motor render html
func (h *Html) SetPage(p *page){
  h.Page= p
}
// render element for html
func (h *Html) RenderElement(e Element)string{
  var value string
  if e.Args().State != nil { 
    e.Args().State.Add(e)
    value= strings.ReplaceAll(e.Args().Value,"{{."+e.Args().State.name+"}}",fmt.Sprint(e.Args().State.value))
  }
  res := fmt.Sprint("<",e.Tag(),e.Args().ToHtml(),">",value)
  for _,item := range e.Children(){
    item.SetMotorRender(h)
    res += item.render()
  }
  res += fmt.Sprint("</",e.Tag(),">")
  return res
}
// render element for html
func (h *Html) RenderPage()(res string){
  var body string
  var headers string

  res= fmt.Sprint("<html>")
  for _,item:= range h.Page.Children(){
    
    if item.Tag() != "style" && item.Tag() != "header"{
      body+= item.render()
    }else{
      headers+= item.render()
    }
  }
  res+= fmt.Sprint(headers,"<body>",body,"</body></html>")
  return res
}
