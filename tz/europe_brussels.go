// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeBrusselsLocation  sync.Once
	cacheEuropeBrusselsLocation *time.Location
)

type EuropeBrussels struct{}

func (EuropeBrussels) Location() *time.Location {
	onceEuropeBrusselsLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Brussels")
		if err != nil {
			panic(err)
		}
		cacheEuropeBrusselsLocation = loc
	})
	return cacheEuropeBrusselsLocation
}
