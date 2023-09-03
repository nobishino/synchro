// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceIndianKerguelenLocation  sync.Once
	cacheIndianKerguelenLocation *time.Location
)

type IndianKerguelen struct{}

func (IndianKerguelen) Location() *time.Location {
	onceIndianKerguelenLocation.Do(func() {
		loc, err := time.LoadLocation("Indian/Kerguelen")
		if err != nil {
			panic(err)
		}
		cacheIndianKerguelenLocation = loc
	})
	return cacheIndianKerguelenLocation
}
