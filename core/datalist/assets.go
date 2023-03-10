package datalist

import (
	"app/core"
	. "app/core"
)

// generate option for tag select html
func generateOptions(childs []string) (res []Element) {
	for _, item := range childs {
		res = append(res, NewElement("item", "option", Args{Value: item}))
	}
	return
}

func Change(e *core.Event) {
	this := e.Target()
	println(this.GetValue())

}
