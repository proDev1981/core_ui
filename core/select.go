package core

func Select(args Args, childs ...string) Element {
	return NewElement(
		"select",
		"select",
		args,
		generateOptions(childs)...)
}
