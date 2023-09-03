// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAsiaVladivostokLocation  sync.Once
	cacheAsiaVladivostokLocation *time.Location
)

type AsiaVladivostok struct{}

func (AsiaVladivostok) Location() *time.Location {
	onceAsiaVladivostokLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Vladivostok")
		if err != nil {
			panic(err)
		}
		cacheAsiaVladivostokLocation = loc
	})
	return cacheAsiaVladivostokLocation
}
