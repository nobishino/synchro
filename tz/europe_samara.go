// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeSamaraLocation  sync.Once
	cacheEuropeSamaraLocation *time.Location
)

type EuropeSamara struct{}

func (EuropeSamara) Location() *time.Location {
	onceEuropeSamaraLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Samara")
		if err != nil {
			panic(err)
		}
		cacheEuropeSamaraLocation = loc
	})
	return cacheEuropeSamaraLocation
}
