package core

import (
	"assets"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode/utf8"
)

// asset compose message type bind for send to client
func ComposeBind(name string) string {
	return `{"type":"bind","name":"` + name + `"}`
}

// asset compose message type event listener for send to client
func ComposeEventListener(id, types, name string) string {
	return `{"type":"eval","js":"document.getElementById('` + id + `').addEventListener('` + types + `',` + name + `)"}`
}

// asset compose message type eval for send to client
func ComposeEval(js string, args ...any) (res string) {
	js = fmt.Sprintf(js, args...)
	res = fmt.Sprintf(`{"type":"eval","js":"%s"}`, js)
	assets.If(!utf8.ValidString(res), "", res)
	return
}

// asset compose message type eval for send to client witch response
func ComposeEvalAnsResponse(name string, js string, args ...any) (res string) {
	js = fmt.Sprintf(js, args...)
	res = fmt.Sprintf(`{"type":"eval","name":"`+name+`","js":"%s"}`, js)
	return
}

// asset goruitne action
func Await(f func()) {
	go f()
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

// assets search
func searchAll(query string, parent Element) (res []Element) {
	if query[0] == '#' {
		if parent.Args().id == query[1:] {
			res = append(res, parent)
		}
	}
	if query[0] == '.' {
		if parent.Args().Class == query[1:] {
			res = append(res, parent)
		}
	}
	if query[0] != '.' && query[0] != '#' {
		if parent.Tag() == query {
			res = append(res, parent)
		}
	}
	for _, item := range parent.Children() {
		res = append(res, searchAll(query, item)...)
	}
	return
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

// return string from file path
func getFile(path string) string {
	return string(assets.Try(os.ReadFile(path)))
}

// destructuring map[string]any
func UnPack[B any](data map[string]B, res ...*B) {
	count := 0
	for key := range data {
		*res[count] = data[key]
		count++
	}
}

// convert type
func ConvertTo[T any](value any) T {
	return value.(T)
}
