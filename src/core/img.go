package core

type img struct {
	Element
}

func Img(args Args) *img {
	return &img{NewElement("img", "img", args)}
}
