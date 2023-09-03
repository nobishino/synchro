// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAsiaKamchatkaLocation  sync.Once
	cacheAsiaKamchatkaLocation *time.Location
)

type AsiaKamchatka struct{}

func (AsiaKamchatka) Location() *time.Location {
	onceAsiaKamchatkaLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Kamchatka")
		if err != nil {
			panic(err)
		}
		cacheAsiaKamchatkaLocation = loc
	})
	return cacheAsiaKamchatkaLocation
}
