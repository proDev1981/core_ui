package core

import "fmt"
import "strings"
import "log"
import "reflect"

// init
func init(){
  log.SetFlags(log.Lshortfile)
}
type Html struct{}
// contructor
func NewHtml()*Html{
  return &Html{}
}
func HtmlRender(p *page)string{
  p.SetMotorRender(NewHtml())
  return p.Render()
}
// render map element
func (h *Html) RenderMap(e Element)string{
  var body string
  var header string
  var footer string
  
  header+= fmt.Sprint("<",e.Tag(),e.Args().ToHtml(),">")
  footer+= fmt.Sprint("</",e.Tag(),">")

  if e.Args().State != nil {
    types:= fmt.Sprintf("%T",e.Args().State)
    if types == "*core.State"{
      data:= e.Args().State.Get()
      Dvalue:= reflect.ValueOf(data)
      Len:= Dvalue.Len()

      for i:=0 ; i<Len ; i++{
        Dstruct:= Dvalue.Index(i)
        body+= fmt.Sprint("<",e.Tag(),e.Args().ToHtml(),">")
        for index,item:= range e.Children(){
          value:= fmt.Sprint(Dstruct.Field(index))
          key:= fmt.Sprint(Dstruct.Type().Field(index).Name)
          body+= fmt.Sprint("<",item.Tag(),item.Args().ToHtml(),">",
                              strings.ReplaceAll(item.Args().Value,"{{."+key+"}}",value),
                            "</",item.Tag(),">")
        }
        body+= fmt.Sprint("</",e.Tag(),">")
      }
      
    }
  }
  return fmt.Sprint(header,body,footer)
}
// render element for html
func (h *Html) RenderElement(e Element)(res string){
  if e.GetSubType() == "map" { 
    e.Args().State.Add(e)
    res= h.RenderMap(e) 
  }else{
    var value string
    if e.Args().State != nil { 
      e.Args().State.Add(e)
      value= strings.ReplaceAll(e.Args().Value,"{{."+e.Args().State.name+"}}",fmt.Sprint(e.Args().State.value))
    }
    res= fmt.Sprint("<",e.Tag(),e.Args().ToHtml(),">",value)
    for _,item:= range e.Children(){
      item.SetMotorRender(h)
      res+= item.render()
    }
    res+= fmt.Sprint("</",e.Tag(),">")
  } 
  return
}
// render element for html
func (h *Html) RenderPage(p *page)(res string){
  var body string
  var headers string

  res= fmt.Sprint("<html>")
  for _,item:= range p.Children(){
    item.SetMotorRender(p.motorRender)
    if item.Tag() != "style" && item.Tag() != "header"{
      body+= item.render()
    }else{
      headers+= item.render()
    }
  }
  res+= fmt.Sprint(headers,"<body>",body,"</body></html>")
  return res
}
