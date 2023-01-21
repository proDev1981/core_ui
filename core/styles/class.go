package styles

import "app/core"
import "fmt"

var attributes = map[string]string{
	"BgColor":        "background-color",
	"BorderRadius":   "border-radius",
	"Color":          "color",
	"Margin":         "margin",
	"BoxShadow":      "box-shadow",
	"Border":         "border",
	"Hover":          "hover",
	"Padding":        "padding",
	"MinHeight":      "min-height",
	"MinWidth":       "min-width",
	"MaxHeight":      "max-height",
	"MaxWidth":       "max-width",
	"Display":        "display",
	"Height":         "height",
	"Width":          "width",
	"FlexDirection":  "flex-direction",
	"JustifyContent": "justify-content",
}

type Attrs struct {
	BgColor        string
	Color          string
	Margin         string
	Padding        string
	Border         string
	BorderRadius   string
	BoxShadow      string
	Hover          string
	MinWidth       string
	MinHeight      string
	MaxWidth       string
	MaxHeight      string
	Display        string
	Height         string
	Width          string
	FlexDirection  string
	JustifyContent string
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
func RGBA(r, g, b, a float32) string {
	return fmt.Sprint(
		"rgba(",
		fmt.Sprint(r), ",",
		fmt.Sprint(g), ",",
		fmt.Sprint(b), ",",
		fmt.Sprint(a), ")",
	)
}
