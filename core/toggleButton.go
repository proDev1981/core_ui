package core

func ToggleButton(args Args, callback func(e *Event, state bool)) Element {
	toggle := false
	if args.Events == nil {
		args.Events = Listener{}
	}
	args.Events["click"] = func(e *Event) {
		toggle = !toggle
		this := e.Target()
		this.SetAttribute("data-toggle", String(toggle))
		callback(e, toggle)
	}
	return Button(args)
}
