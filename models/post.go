package models

type Post struct {
	Email   string
	Content string
}

var Posts = []Post{
	{
		Email:   "a.p.g.Valencia@gmail.com",
		Content: "Tu comentario no me parece adecuado , deberias de pedir perdon.",
	},
	{
		Email:   "realpro.tecnico@gmail.com",
		Content: "Va pedir perdon tu puta madre.",
	},
}
