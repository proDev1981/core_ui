package controler

import "app/core/gcss"

func stylesControler() string {
	return gcss.Sheet{

		"$": &gcss.Rule{
			Display: "flex",
		},
		"$button": &gcss.Rule{
			Color:   "blue",
			BgColor: "white",
			Border:  "1px blue solid",
			Hover: &gcss.Rule{
				BgColor: "blue",
				Color:   "white",
			},
		},
	}.Parse()
}
