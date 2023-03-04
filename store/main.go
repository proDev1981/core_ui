package store

import "app/core"
import "app/models"

var States = core.Provider{
	"data": core.NewState(models.Posts),
}
