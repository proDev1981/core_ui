package core

func DataList(args Args, childs ...string) Element {
	datalist := NewElement(
		"datalist",
		"datalist",
		Args{},
		generateOptions(childs)...,
	)
	return NewElement(
		"datalist-container",
		"div",
		args,
		Input(Args{List: datalist.GetId()}),
		datalist,
	)
}
