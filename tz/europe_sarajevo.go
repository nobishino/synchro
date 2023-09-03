// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeSarajevoLocation  sync.Once
	cacheEuropeSarajevoLocation *time.Location
)

type EuropeSarajevo struct{}

func (EuropeSarajevo) Location() *time.Location {
	onceEuropeSarajevoLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Sarajevo")
		if err != nil {
			panic(err)
		}
		cacheEuropeSarajevoLocation = loc
	})
	return cacheEuropeSarajevoLocation
}
