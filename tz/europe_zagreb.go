// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeZagrebLocation  sync.Once
	cacheEuropeZagrebLocation *time.Location
)

type EuropeZagreb struct{}

func (EuropeZagreb) Location() *time.Location {
	onceEuropeZagrebLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Zagreb")
		if err != nil {
			panic(err)
		}
		cacheEuropeZagrebLocation = loc
	})
	return cacheEuropeZagrebLocation
}
