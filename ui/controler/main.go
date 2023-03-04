package controler

import . "app/core"

func New() Element {
	return Box(Args{
		Key: "controler",
	},
		Button(Args{
			Value:  "get value",
			Events: Listener{"click": change},
		}),
		Styles(stylesControler()),
	)
}
