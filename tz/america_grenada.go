// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAmericaGrenadaLocation  sync.Once
	cacheAmericaGrenadaLocation *time.Location
)

type AmericaGrenada struct{}

func (AmericaGrenada) Location() *time.Location {
	onceAmericaGrenadaLocation.Do(func() {
		loc, err := time.LoadLocation("America/Grenada")
		if err != nil {
			panic(err)
		}
		cacheAmericaGrenadaLocation = loc
	})
	return cacheAmericaGrenadaLocation
}
