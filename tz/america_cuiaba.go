// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAmericaCuiabaLocation  sync.Once
	cacheAmericaCuiabaLocation *time.Location
)

type AmericaCuiaba struct{}

func (AmericaCuiaba) Location() *time.Location {
	onceAmericaCuiabaLocation.Do(func() {
		loc, err := time.LoadLocation("America/Cuiaba")
		if err != nil {
			panic(err)
		}
		cacheAmericaCuiabaLocation = loc
	})
	return cacheAmericaCuiabaLocation
}
