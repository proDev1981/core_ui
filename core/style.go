package core

// make element with gcss
func Styles(css string) *Ele {
	return &Ele{
		tag:  "style",
		args: Args{Value: css},
	}
}

// get css from file css with path and make element
func StylesFrom(path string) *Ele {
	return &Ele{
		tag:  "style",
		args: Args{Value: Mimi(getFile(path))},
	}
}
