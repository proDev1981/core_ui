package core

import (
	"strings"
)

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
	"AlignItems":     "align-items",
	"BoxSizing":      "box-sizing",
	"OverFlow":       "overflow",
	"Transition":     "transition",
}

type Sheet map[string]*Rule

func (s Sheet) Parse() (res string) {
	for key, value := range s {
		if strings.Contains(key, "global") {
			key = strings.Replace(key, "global ", "", 1)
		} else {
			key = "$" + key
		}
		res += value.Parse(key)
	}
	return
}
func ToCss(sheet Sheet) string {
	return sheet.Parse()
}

type Rule struct {
	BgColor        string
	Color          string
	Margin         string
	Padding        string
	Height         string
	Width          string
	MinHeight      string
	MinWidth       string
	MaxHeight      string
	MaxWidth       string
	Border         string
	BorderColor    string
	BorderRadius   string
	BoxShadow      string
	BoxSizing      string
	Display        string
	JustifyContent string
	AlignItems     string
	FlexDirection  string
	OverFlow       string
	Transition     string
	Hover          *Rule
}

// parse gcss to css vanilla
func (r *Rule) Parse(name string) (res string) {
	return name + parse(name, r)
}

// function parse gcss to css vanilla
func parse(name string, attrs *Rule) (res string) {
	var hover string
	res += "{"
	a := Entries(*attrs)
	for key, value := range a {
		if value != "" {
			if key == "Hover" && String(attrs.Hover) != "<nil>" {
				hover = name + ":hover" + parse(name, attrs.Hover)
			} else if key != "Hover" {
				res += (attributes[key] + ":" + value + ";")
			}
		}
	}
	return res + "}" + hover
}
