package inputs

import . "app/core"

func New() Element {
	return Box(Args{
		Key: "container-inputs",
	},
		Input(Args{
			Placeholder: "background",
		}),
		Input(Args{
			Placeholder: "title",
		}),
	)
}
