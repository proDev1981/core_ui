package core

type container struct{
  Element
}
func Row(args Args,children ...Element)*container{
  r := &container{ NewElement("div",Args{}) }
  r.childs(children...)  
  return r
}
func Column(args Args,children ...Element)*container{
  r := &container{ NewElement("div",Args{}) }
  r.childs(children...)  
  return r
}
