// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	oncePacificGambierLocation  sync.Once
	cachePacificGambierLocation *time.Location
)

type PacificGambier struct{}

func (PacificGambier) Location() *time.Location {
	oncePacificGambierLocation.Do(func() {
		loc, err := time.LoadLocation("Pacific/Gambier")
		if err != nil {
			panic(err)
		}
		cachePacificGambierLocation = loc
	})
	return cachePacificGambierLocation
}
