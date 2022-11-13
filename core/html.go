package core

import "fmt"
import "strings"
import "log"
import "reflect"
import "os"
import "app/core/http"

var dom *page

// init
func init() {
	log.SetFlags(log.Lshortfile)
}

type Html struct {
	server *http.Server
	dom    *page
}

// contructor
func NewHtml() *Html {
	return &Html{server: http.NewServer()}
}

// builder page html
func HtmlBuild(p *page) Motor {
	p.SetMotorRender(NewHtml())
	if _, err := os.Stat("./src"); err != nil {
		log.Println("No exist folder => src ")
		log.Println("Creating this..")
		if err = os.Mkdir("./src", os.ModePerm); err != nil {
			panic(err)
		}
	}
	file, err := os.Create("./src/index.html")
	if err != nil {
		panic(err)
	}
	file.Write([]byte(p.Render()))
	p.motorRender.(*Html).dom = p
	dom = p
	return p.motorRender
}

// add listener in struct html
// @ params
// parent is id of element html
// listener is map[type evenet]function event
func (h *Html) AddEventListener(parent string, listener Listener) {
	for types, call := range listener {
		if h.server.GetSocket() == nil {
			h.server.SetInitialEvents(parent, types, call)
		} else {
			name := types + parent
			h.server.GetSocket().SendAll(http.ComposeBind(name))
			h.server.GetSocket().SendAll(http.ComposeEventListener(parent, types, name))
		}
	}
}

// getter server
func (h *Html) NewServer() *http.Server {
	return h.server
}

// render map element
func (h *Html) RenderMap(e Element) string {
	var body string
	var header string
	var footer string

	header += fmt.Sprint("<", e.Tag(), argsToHTml(e.Args()), ">")
	footer += fmt.Sprint("</", e.Tag(), ">")

	if e.Args().State != nil {
		types := fmt.Sprintf("%T", e.Args().State)
		if types == "*core.State" {
			data := e.Args().State.Get()
			Dvalue := reflect.ValueOf(data)
			Len := Dvalue.Len()

			for i := 0; i < Len; i++ {
				Dstruct := Dvalue.Index(i)
				body += fmt.Sprint("<", e.Tag(), argsToHTml(e.Args()), ">")
				for index, item := range e.Children() {
					item.setParent(e)
					value := fmt.Sprint(Dstruct.Field(index))
					key := fmt.Sprint(Dstruct.Type().Field(index).Name)
					body += fmt.Sprint("<", item.Tag(), argsToHTml(item.Args()), ">",
						strings.ReplaceAll(item.Args().Value, "{{."+key+"}}", value),
						"</", item.Tag(), ">")
				}
				body += fmt.Sprint("</", e.Tag(), ">")
			}

		}
	}
	return fmt.Sprint(header, body, footer)
}

// render element for html
func (h *Html) RenderElement(e Element) (res string) {
	if e.GetSubType() == "map" {
		e.Args().State.Add(e)
		res = h.RenderMap(e)
	} else {
		//
		log.Println(e.Tag())
		log.Println(e.Children())
		//
		var value string
		if e.Args().State != nil {
			e.Args().State.Add(e)
			value = replaceState(e.Args(), "")
		} else {
			value = e.Args().Value
		}
		res = fmt.Sprint("<", e.Tag(), argsToHTml(e.Args()), ">", value)
		for _, item := range e.Children() {
			item.setParent(e)
			item.SetMotorRender(h)
			res += item.render()
			if !eventsIsEmpty(item.Args().Events) {
				h.AddEventListener(item.Args().id, item.Args().Events)
			}
		}
		res += fmt.Sprint("</", e.Tag(), ">")
	}
	return
}

// render element for html
func (h *Html) RenderPage(p *page) (res string) {
	var body string
	var headers string

	res = fmt.Sprint("<html>")
	for _, item := range p.Children() {
		item.SetMotorRender(p.motorRender)
		if item.Tag() != "style" && item.Tag() != "header" {
			body += item.render()
		} else {
			headers += item.render()
		}
	}
	jsScript := "<script src='index.js'></script>"
	res += fmt.Sprint(headers, "<body>", jsScript, body, "</body></html>")
	return res
}

// search element by query
func (h *Html) Selector(query string) Element {
	if h.dom != nil {
		return search(query, h.dom.children[0])
	}
	return nil
}
func Selector(query string) Element {
	if dom != nil {
		return search(query, dom.children[0])
	}
	return nil

}

// assets search
func search(query string, parent Element) Element {
	if query[0] == '#' {
		if parent.Args().id == query[1:] {
			return parent
		}
	}
	if query[0] == '.' {
		if parent.Args().Class == query[1:] {
			return parent
		}
	}
	for _, item := range parent.Children() {
		if res := search(query, item); res != nil {
			return res
		}
	}
	return nil
}

// obtengo si la interface en un struct
func isStruct(v any) bool {
	return fmt.Sprint(reflect.TypeOf(v).Kind()) == "struct"
}

// obtengo si la interface es un slice
func isSlice(v any) bool {
	return fmt.Sprint(reflect.TypeOf(v).Kind()) == "slice"
}

// convert struct[any] to map[key struct]value struct
func entries(v any) map[string]string {
	res := make(map[string]string)
	T := reflect.TypeOf(v)
	V := reflect.ValueOf(v)
	Len := T.NumField()
	for i := 0; i < Len; i++ {
		key := T.Field(i).Name
		value := V.Field(i).String()
		res[key] = value
	}
	return res
}

// replace args.value for state.value
func replaceState(e Args, change string) (value string) {

	if len(change) > 0 {
		value = change
	} else {
		value = e.Value
	}

	if isStruct(e.State.Get()) {
		// if value state is struct[any]
		for key, val := range entries(e.State.Get()) {
			value = strings.ReplaceAll(value, "{{."+key+"}}", val)
		}
		// if value state is []any
	} else if isSlice(e.State.Get()) {
		// if value struct is string|int|float
	} else {
		value = strings.ReplaceAll(value, "{{.state}}", fmt.Sprint(e.State.value))
	}
	return
}

// args to html parse
func argsToHTml(s Args) string {
	var res string
	sType := reflect.TypeOf(s)
	sValue := reflect.ValueOf(s)
	sLen := sType.NumField()

	for i := 0; i < sLen; i++ {

		value := fmt.Sprint(sValue.Field(i))
		name := strings.ToLower(sType.Field(i).Name)
		types := fmt.Sprint(sType.Field(i).Type)

		if types == "string" && name != "direction" && name != "value" && len(value) > 0 {
			if s.State != nil {
				value = replaceState(s, value)
			}
			res += fmt.Sprint(" ", name, "='", value, "'")
		}
	}
	return res
}

// assets
func eventsIsEmpty(e Listener) bool {
	return len(e) <= 0
}
