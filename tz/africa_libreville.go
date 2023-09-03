// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAfricaLibrevilleLocation  sync.Once
	cacheAfricaLibrevilleLocation *time.Location
)

type AfricaLibreville struct{}

func (AfricaLibreville) Location() *time.Location {
	onceAfricaLibrevilleLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Libreville")
		if err != nil {
			panic(err)
		}
		cacheAfricaLibrevilleLocation = loc
	})
	return cacheAfricaLibrevilleLocation
}
