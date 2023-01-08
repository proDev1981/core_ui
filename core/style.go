package core

// falta por implementar
func Styles(args Args) *Ele {

	if args.Src != "" {
		args.Value = Mimi(getFile(args.Src)) + args.Value
	}
	return &Ele{
		tag:  "style",
		args: args,
	}
}
