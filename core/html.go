package core

import "fmt"
import "regexp"
import "strings"
import "log"
import "reflect"
import "os"
import "assets"
import "extend/meta"
import "encoding/json"
import "github.com/gorilla/websocket"

var dom Element

// init
func init() {
	log.SetFlags(log.Lshortfile)
}

type Html struct {
	server   *Server
	dom      Element
	conn     *websocket.Conn
	elements map[string]Element
}

var html *Html

// contructor
func NewHtml() *Html {
	if html == nil {
		html = &Html{
			server:   NewServer(),
			elements: make(map[string]Element),
		}
	}
	return html
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
	return p.motorRender
}

// add listener in struct html
// @ params
// parent is id of element html
// listener is [type evenet]function event
func (h *Html) AddEventListener(parent string, listener Listener) {
	for types, call := range listener {
		if h.server.GetSocket() == nil {
			h.server.SetInitialEvents(parent, types, call)
		} else {
			name := types + parent
			h.server.GetSocket().SendAll(ComposeBind(name))
			h.server.GetSocket().SendAll(ComposeEventListener(parent, types, name))
		}
	}
}

// getter server
func (h *Html) NewServer() *Server {
	return h.server
}

// getter server
func (h *Html) GetServer() *Server {
	return h.server
}

// render for
func (h *Html) renderFor(e Element) (res string) {
	return
}

// replace attribute each if is reactive
func replaceAttribute(e Element, str string) (res string) {
	res = str
	if strings.Contains(str, "{{") && strings.Contains(str, "}}") {
		name := strings.Replace(strings.Replace(str, "{{", "", 1), "}}", "", 1)
		var value string
		state := Store[name]
		if state != nil {
			state.Fill(&value)
			state.AddElement(e)
			res = strings.Replace(str, "{{"+name+"}}", value, 1)
		}
	}
	return
}

// render nuevo
func (h *Html) RenderMap(e Element) string {
	header := fmt.Sprint("<", e.Tag(), argsToHTml(e.Args()), ">")
	each := replaceAttribute(e, e.Args().Each)
	max := replaceAttribute(e, e.Args().Max)
	log.Println("each:", each)
	var body string
	var state *State
	if each != "" {
		state = Store[each]
		if state != nil {
			state.AddElement(e)
			// get data with reflect
			r_data := reflect.ValueOf(state.Get())
			if String(r_data.Kind()) == "slice" {
				len_data := r_data.Len()
				if len_data > Int(max) && Int(max) > 0 {
					len_data = Int(max)
				}
				if len_data > 0 {
					for i := 0; i < len_data; i++ {
						Struct := r_data.Index(i)
						// loop children
						for _, item := range e.Children() {
							item.setId("") //delete id in element mapper
							item.setParent(e)
							item.SetMotorRender(h)
							inner := item.render()
							// replace all values states
							for index := 0; index < Struct.NumField(); index++ {
								name := "{{" + Struct.Type().Field(index).Name + "}}"
								value := Struct.Field(index)
								inner = strings.ReplaceAll(inner, name, fmt.Sprint(value))
							}
							body += inner
						}
					}

				}
			}
		}
	}

	foother := fmt.Sprint("</", e.Tag(), ">")
	return header + body + foother
}

// new render Element
func (h *Html) RenderElement(e Element) (res string) {
	// compruebo si el elemento existe
	conditional := func() bool { return true }
	if e.Args().Case != nil {
		conditional = e.Args().Case
	}
	log.Println("case: ", conditional(), "tag: ", e.Tag())
	if h.saveKey(e) && conditional() {
		// parseo los estilos
		h.parseStyle(e)
		// compruebo si es un elemento list
		switch e.GetSubType() {
		case "list":
			res = h.RenderMap(e)
		default:
			// convierto elemento en string
			res = toString(h, e)
			// compruebo que el elemento es reactivo o no es un list
			if e.IsReactive() {
				// remplazo valor del stato en el elemento
				res = replaceAllSates(e, res)
			}
		}
	}
	return
}

/* ASSETS */

// replace attribute if is reactive
func replaceMax(e Element) (res string) {
	str := e.Args().Each
	res = str
	if strings.Contains(str, "{{") && strings.Contains(str, "}}") {
		name := strings.Replace(strings.Replace(str, "{{", "", 1), "}}", "", 1)
		var value string
		state := Store[name]
		if state != nil {
			state.Fill(&value)
			state.AddElement(e)
			res = strings.Replace(str, "{{"+name+"}}", value, 1)
		}
	}
	return
}

// paso de struct element o string
func toString(h *Html, e Element) (res string) {
	attributes := argsToHTml(e.Args())
	inner := e.Args().Value
	// render childrens
	for _, item := range e.Children() {
		item.setParent(e)
		item.SetMotorRender(h)
		inner += item.render()
		if !eventsIsEmpty(item.Args().Events) {
			h.AddEventListener(item.Args().id, item.Args().Events)
		}
	}
	res = fmt.Sprintf("<%s %s>%s</%s>", e.Tag(), attributes, inner, e.Tag())
	return
}

// replace states in htmls string
func replaceAllSates(e Element, res string) string {
	// compilo la expresion regular
	r, err := regexp.Compile(`\{\{[\w\d\.]+\}\}`)
	switch err {
	case nil:
		var state *State
		var val string
		matches := r.FindAllString(res, -1)
		for _, match := range matches {
			// limpio la varible del template
			name := strings.ReplaceAll(strings.ReplaceAll(match, "{", ""), "}", "")
			// compruebo si esta pidiendo un objeto o un valor primitivo
			sliceMatch := TrimSlice(strings.Split(name, "."), "")
			switch {
			case len(sliceMatch) <= 1: // value primitive
				name = sliceMatch[0]
				// rescato stado del store
				if Store[name] != nil {
					state = Store[name]
					state.AddElement(e)
					val = state.ToString()
					// remplazar varibles con su valor en el estado
					res = strings.ReplaceAll(res, match, val)
				}
			case len(sliceMatch) > 1: // value complex
				name = sliceMatch[0]
				if Store[name] != nil {
					state = Store[name]
					state.AddElement(e)
					val = meta.Entries(state.Get())[sliceMatch[1]]
					// remplazar varibles con su valor en el estado
					res = strings.ReplaceAll(res, match, val)
				}
			}
		}
	default:
		fmt.Println("Error:", err.Error())
	}
	return res
}

/* END ASSETS */

// render element for html
func (h *Html) RenderPage(p *page) (res string) {
	var body string
	var headers string

	res = fmt.Sprint("<!DOCTYPEhtml><html lang='" + p.lang + "'>")
	for _, item := range p.Children() {
		item.SetMotorRender(p.motorRender)
		if item.Tag() != "style" && item.Tag() != "head" {
			if item.Args().Name == "root" {
				dom = item
				h.dom = item
			}
			body += item.render()
		} else {
			headers += item.render()
		}
	}
	res += fmt.Sprint(headers, "<body>", body, "</body></html>")
	return res
}

// search element by query
func (h *Html) RootSelector(query string) Element {
	return selector(dom, query)
}

// search element in map element by key
func (h *Html) Key(key string) Element {
	ele := h.elements[key]
	if ele == nil {
		panic(fmt.Sprintf("<Element key:'%s'> not found!!", key))
	}
	return ele
}

// search element by query
func (h *Html) Selector(e Element, query string) Element {
	return selector(e, query)
}

// search element by query
func (h *Html) SelectorAll(e Element, query string) []Element {
	return selectorAll(e, query)
}

// search element by query
func selector(e Element, query string) Element {
	if e != nil {
		return search(query, e)
	}
	return nil

}

// search element by query
func selectorAll(e Element, query string) (res []Element) {
	if e != nil {
		res = append(res, searchAll(query, e)...)
	}
	return

}

// setter conn
func (h *Html) SetConn(conn *websocket.Conn) {
	h.conn = conn
}

// getter conn
func (h *Html) Conn() *websocket.Conn {
	return h.conn
}

// update element in dom of client
func (h *Html) Update(e Element) {
	compose := ComposeEval("document.getElementById('%s').outerHTML =`%s`",
		e.Args().id, e.render(),
	)
	h.Conn().WriteMessage(1, []byte(compose))
}

// create new objetc in js
func (h *Html) NewObject(e Element, name string, value any) string {
	p := NewPromise(h)
	val := strings.ReplaceAll(string(assets.Try(json.Marshal(value))), `"`, "'")
	createObject := ComposeEval(`window.%s = new Object`, name)
	fillObject := ComposeEvalAnsResponse(p.id, `%s=%s`, name, val)
	h.Conn().WriteMessage(1, []byte(createObject))
	h.Conn().WriteMessage(1, []byte(fillObject))
	return <-p.Await()
}

// set attribute of element in dom html
func (h *Html) SetAttribute(e Element, name, value string) {
	var cmd string
	if name == "value" {
		if e.GetSubType() == "button" {
			name = "innerHTML"
		}
		cmd = `document.getElementById('%s').%s='%s'`
	} else {
		cmd = `document.getElementById('%s').setAttribute('%s','%s')`
	}
	//p := NewPromise(h)
	setAttribute := ComposeEval(
		cmd,
		e.Args().id,
		name,
		value,
	)
	h.Conn().WriteMessage(1, []byte(setAttribute))
	//return <-p.Await()
}

// geter attribute of element in dom html
func (h *Html) GetAttribute(e Element, name string) string {
	p := NewPromise(h)
	var getAttribute string
	if name == "value" {
		if e.GetSubType() == "button" {
			getAttribute = ComposeEvalAnsResponse(p.id,
				`document.getElementById('%s').innerHTML`,
				e.Args().id,
			)
		} else {
			getAttribute = ComposeEvalAnsResponse(p.id,
				`document.getElementById('%s').value`,
				e.Args().id,
			)
		}
	} else {
		getAttribute = ComposeEvalAnsResponse(p.id,
			`document.getElementById('%s').getAttribute('%s')`,
			e.Args().id,
			name,
		)
	}
	h.Conn().WriteMessage(1, []byte(getAttribute))
	return <-p.Await()
}

// log string in js console
func (h *Html) Log(e Element, value ...string) {
	log := ComposeEval(`console.log('%s')`, value)
	h.Conn().WriteMessage(1, []byte(log))
}

// alert in js client
func (h *Html) Alert(e Element, value ...string) {
	alert := ComposeEval(`alert('%s')`, value)
	h.Conn().WriteMessage(1, []byte(alert))
}

// get target
func Target(e *Event) Element {
	ele := selector(dom, "#"+e.Id)
	ele.MotorRender().SetConn(e.Client)
	return ele
}

// return map string string values input in element
func (h *Html) GetData(e Element) map[string]string {

	children := e.SelectorAll("input")
	patter := make(map[string]string)
	if len(children) > 0 {
		for index, item := range children {
			switch {
			case item.Args().Placeholder != "":
				patter[item.Args().Placeholder] = item.GetAttribute("value")
			case item.Args().Key != "":
				patter[item.Args().Key] = item.GetAttribute("value")
			case item.Args().Class != "":
				patter[item.Args().Class] = item.GetAttribute("value")
			case item.Args().Name != "":
				patter[item.Args().Name] = item.GetAttribute("value")
			default:
				patter[String(index)] = item.GetAttribute("value")
			}
		}
	}
	return patter
}

// delete value element
func (h *Html) Reset(e Element) Element {
	e.SetAttribute("value", "")
	return e
}

// focus element
func (h *Html) Focus(e Element) Element {
	focus := ComposeEval(`document.getElementById('%s').focus()`, e.Args().id)
	h.Conn().WriteMessage(1, []byte(focus))
	return e
}

// get innerhtml element
func (h *Html) GetInner(e Element) string {
	p := NewPromise(h)
	getInnerHtml := ComposeEvalAnsResponse(p.id,
		"document.getElementById('%s').innerHTML",
		e.Args().id,
	)
	h.Conn().WriteMessage(1, []byte(getInnerHtml))
	return <-p.Await()
}

// set innerHTML element
func (h *Html) SetInner(e Element, value string) {
	setInnerHTml := ComposeEval(
		"document.getElementById('%s').innerHTML = '%s'",
		e.Args().id,
		value,
	)
	h.Conn().WriteMessage(1, []byte(setInnerHTml))
}

// get value element
func (h *Html) GetValue(e Element) string {
	return h.GetAttribute(e, "value")
}

// set value element
func (h *Html) SetValue(e Element, value string) {
	h.SetAttribute(e, "value", value)
}

// themes
func (h *Html) SetBackgroundColor(color string) {
	setcolor := ComposeEval(
		`document.querySelector('%s').style.backgroundColor='%s'`,
		"body",
		color,
	)
	h.Conn().WriteMessage(1, []byte(setcolor))
}

func (h *Html) SetBackgroundTitle(color string) {
	setcolor := ComposeEval(
		`document.querySelector('%s').content='%s'`,
		"#color-title",
		color,
	)
	h.Conn().WriteMessage(1, []byte(setcolor))
}

/* zone apperence */
// add class in array class element
func (h *Html) AddClass(e Element, value string) {
	addClass := ComposeEval(
		"document.getElementById('%s').classList.add('%s')",
		e.Args().id,
		value,
	)
	h.Conn().WriteMessage(1, []byte(addClass))
}

// add class in array class element
func (h *Html) RemoveClass(e Element, value string) {
	removeClass := ComposeEval(
		"document.getElementById('%s').classList.remove('%s')",
		e.Args().id,
		value,
	)
	h.Conn().WriteMessage(1, []byte(removeClass))
}

/* ASSETS */

// save element in map with hash
func (h *Html) saveKey(e Element) bool {
	key := e.Args().Key
	if key != "" {
		if h.elements[key] != nil {
			log.Printf("Warning: Element <<%s>> duplicate!!\n", key)
		}
		h.elements[key] = e
	}
	return true
}

// replace hash in css
func (h *Html) parseStyle(e Element) {
	if e.Tag() == "style" {
		parent := fmt.Sprint("[ id = '", e.Parent().Args().id, "'] ")
		e.SetArgs(Args{Value: strings.ReplaceAll(e.Args().Value, "$", parent)})

	}
}

// replace all item match
func TrimSlice(s []string, query string) (res []string) {
	for _, item := range s {
		if item != query {
			res = append(res, item)
		}
	}
	return
}
