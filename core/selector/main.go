package selector

import . "app/core"

func New(args Args, childs ...string) Element {
	return NewElement(
		"selector",
		"select",
		args,
		generateOptions(childs)...)
}
