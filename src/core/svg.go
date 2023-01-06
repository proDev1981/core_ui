package core

import (
	"assets"
	"regexp"
	"strings"
)

type svg struct {
	Element
}

func Svg(args Args) *svg {
	cnt := getFile(args.Src)
	start := strings.Index(cnt, ">") + 1
	header := cnt[:start]
	cnt = cnt[start:]
	cnt = strings.Replace(cnt, "</svg>", "", 1)
	args.Value = cnt
	if args.Width == "" {
		args.Width = getWitdh(header)
	}
	if args.Height == "" {
		args.Height = getHeight(header)
	}
	return &svg{NewElement("svg", "svg", args)}
}

func getWitdh(str string) string {
	return getAttr(str, "width")
}
func getHeight(str string) string {
	return getAttr(str, "height")
}

func getAttr(str string, name string) string {

	reg := assets.Try(regexp.Compile(name + `\s*=\s*"[\d\w]+"\s`))
	all := reg.FindAllString(str, 1)
	if len(all) > 0 {
		return strings.ReplaceAll(strings.Split(all[0], "=")[1], `"`, "")
	}
	return ""
}
