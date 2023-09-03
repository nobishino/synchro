// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	oncePacificNauruLocation  sync.Once
	cachePacificNauruLocation *time.Location
)

type PacificNauru struct{}

func (PacificNauru) Location() *time.Location {
	oncePacificNauruLocation.Do(func() {
		loc, err := time.LoadLocation("Pacific/Nauru")
		if err != nil {
			panic(err)
		}
		cachePacificNauruLocation = loc
	})
	return cachePacificNauruLocation
}
