// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAmericaIndianaTell_CityLocation  sync.Once
	cacheAmericaIndianaTell_CityLocation *time.Location
)

type AmericaIndianaTell_City struct{}

func (AmericaIndianaTell_City) Location() *time.Location {
	onceAmericaIndianaTell_CityLocation.Do(func() {
		loc, err := time.LoadLocation("America/Indiana/Tell_City")
		if err != nil {
			panic(err)
		}
		cacheAmericaIndianaTell_CityLocation = loc
	})
	return cacheAmericaIndianaTell_CityLocation
}
