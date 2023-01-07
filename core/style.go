package core

// falta por implementar
func Styles(path string) *Ele {

	return &Ele{
		tag: "style",
		args: Args{
			Value: getFile(path)},
	}
}
