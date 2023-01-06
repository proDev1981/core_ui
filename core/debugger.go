package core

import "time"
import "log"

func DebugAndRes[T any](sms string, f func() T) (res T) {
	start := time.Now()
	res = f()
	log.Println(sms, time.Now().Sub(start))
	return
}
