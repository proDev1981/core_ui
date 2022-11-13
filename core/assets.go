package core

import "fmt"

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
	return
}
