// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceIndianAntananarivoLocation  sync.Once
	cacheIndianAntananarivoLocation *time.Location
)

type IndianAntananarivo struct{}

func (IndianAntananarivo) Location() *time.Location {
	onceIndianAntananarivoLocation.Do(func() {
		loc, err := time.LoadLocation("Indian/Antananarivo")
		if err != nil {
			panic(err)
		}
		cacheIndianAntananarivoLocation = loc
	})
	return cacheIndianAntananarivoLocation
}
