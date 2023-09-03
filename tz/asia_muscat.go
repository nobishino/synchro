// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAsiaMuscatLocation  sync.Once
	cacheAsiaMuscatLocation *time.Location
)

type AsiaMuscat struct{}

func (AsiaMuscat) Location() *time.Location {
	onceAsiaMuscatLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Muscat")
		if err != nil {
			panic(err)
		}
		cacheAsiaMuscatLocation = loc
	})
	return cacheAsiaMuscatLocation
}
