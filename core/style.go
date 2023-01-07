package core

import "os"
import "log"

// falta por implementar
func Styles(path string) *Ele {

	var data []byte
	data, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error:", err)
	}

	return &Ele{
		tag:  "style",
		args: Args{Value: string(data)},
	}
}
