// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeBerlinLocation  sync.Once
	cacheEuropeBerlinLocation *time.Location
)

type EuropeBerlin struct{}

func (EuropeBerlin) Location() *time.Location {
	onceEuropeBerlinLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Berlin")
		if err != nil {
			panic(err)
		}
		cacheEuropeBerlinLocation = loc
	})
	return cacheEuropeBerlinLocation
}
