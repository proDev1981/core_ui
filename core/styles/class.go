package styles

import "app/core"

var attributes = map[string]string{
	"BgColor":      "background-color",
	"BorderRadius": "border-radius",
	"Color":        "color",
	"Margin":       "margin",
	"BoxShadow":    "box-shadow",
	"Border":       "border",
	"Hover":        "hover",
	"Padding":      "padding",
}

type Attrs struct {
	BgColor      string
	Color        string
	Margin       string
	Padding      string
	Border       string
	BorderRadius string
	BoxShadow    string
	Hover        string
}

func Rule(name string, attrs Attrs) (res string) {
	res += name + "{"
	var hover string
	a := core.Entries(attrs)
	for key, value := range a {
		if value != "" {
			if key == "Hover" {
				hover = name + ":hover" + value
			} else {
				res += (attributes[key] + ":" + value + ";")
			}
		}
	}

	return res + "}" + hover
}
func Hover(attrs Attrs) (res string) {
	res += "{"
	a := core.Entries(attrs)
	for key, value := range a {
		if value != "" {
			res += (attributes[key] + ":" + value + ";")
		}
	}
	return res + "}"
}
