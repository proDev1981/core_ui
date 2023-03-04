package controler

import "app/core"
import "app/store"
import "app/models"

//change
func change(e *core.Event) {
	this := e.Target()
	data := store.States["data"]
	data.Add(models.Post{
		Email:   "nuevo mail",
		Content: "nuevo contenido",
	})
	this.Log(core.String(data.Len()))
	data.Sub(0)
	this.Log(core.String(data.Len()))
}
