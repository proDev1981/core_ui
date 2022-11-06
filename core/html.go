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
  
  header+= fmt.Sprint("<",e.Tag(),argsToHTml(e.Args()),">")
  footer+= fmt.Sprint("</",e.Tag(),">")

  if e.Args().State != nil {
    types:= fmt.Sprintf("%T",e.Args().State)
    if types == "*core.State"{
      data:= e.Args().State.Get()
      Dvalue:= reflect.ValueOf(data)
      Len:= Dvalue.Len()

      for i:=0 ; i<Len ; i++{
        Dstruct:= Dvalue.Index(i)
        body+= fmt.Sprint("<",e.Tag(),argsToHTml(e.Args()),">")
        for index,item:= range e.Children(){
          value:= fmt.Sprint(Dstruct.Field(index))
          key:= fmt.Sprint(Dstruct.Type().Field(index).Name)
          body+= fmt.Sprint("<",item.Tag(),argsToHTml(item.Args()),">",
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
      value= replaceState(e.Args(),"")
    }
    res= fmt.Sprint("<",e.Tag(),argsToHTml(e.Args()),">",value)
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
// obtengo si la interface en un struct
func isStruct(v any)bool{
  return fmt.Sprint(reflect.TypeOf(v).Kind()) == "struct"
}
// obtengo si la interface es un slice
func isSlice(v any)bool{
  return fmt.Sprint(reflect.TypeOf(v).Kind()) == "slice"
}
// struct any to map
func structToMap(v any)map[string]string{
  res:= make(map[string]string)
  T:= reflect.TypeOf(v)
  V:= reflect.ValueOf(v)
  Len:= T.NumField()
  for i:=0 ; i<Len ; i++{
    key:= T.Field(i).Name
    value:= V.Field(i).String()
    res[key]=value
  }
  return res
}
// replace args.value for state.value
func replaceState(e Args,change string)(value string){
  if len(change)>0 { value= change }else{ value= e.Value }
  if isStruct(e.State.Get()){
    log.Println("<< is struct >>")
    for key,val:= range structToMap(e.State.Get()){
      value= strings.ReplaceAll(value,"{{."+key+"}}",val)
    }
    
  }else if isSlice(e.State.Get()){
    log.Println("<< is slice >>")

  }else{
    log.Println("<< is other >>")
    value= strings.ReplaceAll(value,"{{.state}}",fmt.Sprint(e.State.value))
  }
  return
}
// args to html parse
func argsToHTml(s Args)string{
  sType:= reflect.TypeOf(s) 
  sValue:= reflect.ValueOf(s)
  sLen:= sType.NumField()
  var res string

  for i:= 0 ; i<sLen ; i++{
    value:= fmt.Sprint(sValue.Field(i))
    name:= strings.ToLower(sType.Field(i).Name)
    types:= fmt.Sprint(sType.Field(i).Type)
    if s.State != nil {
      if types == "string" && name != "direction" && name != "id" && name != "value" && len(value) > 0{
        value= replaceState(s,value)
        res += fmt.Sprint(" ",name,"='",value,"'")
    }
    }
  }
  return res
}
