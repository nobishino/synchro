// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAfricaBissauLocation  sync.Once
	cacheAfricaBissauLocation *time.Location
)

type AfricaBissau struct{}

func (AfricaBissau) Location() *time.Location {
	onceAfricaBissauLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Bissau")
		if err != nil {
			panic(err)
		}
		cacheAfricaBissauLocation = loc
	})
	return cacheAfricaBissauLocation
}
