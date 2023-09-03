// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeJerseyLocation  sync.Once
	cacheEuropeJerseyLocation *time.Location
)

type EuropeJersey struct{}

func (EuropeJersey) Location() *time.Location {
	onceEuropeJerseyLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Jersey")
		if err != nil {
			panic(err)
		}
		cacheEuropeJerseyLocation = loc
	})
	return cacheEuropeJerseyLocation
}
