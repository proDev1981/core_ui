package color

import "fmt"
import . "app/core"

func RGBA(r, g, b, a float32) string {
	return fmt.Sprint(" rgba(", String(r), ",", String(g), ",", String(b), ",", String(a), ")")
}
