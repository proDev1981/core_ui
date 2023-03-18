package core

type class struct {
	Colapsed         *Rule
	TransitionHeight *Rule
}

var Class = class{
	Colapsed: &Rule{
		Height:   "0",
		OverFlow: "hidden",
		Padding:  "0px",
		Margin:   "0px",
	},
	TransitionHeight: &Rule{
		Transition: "height .3s ease",
	},
}
