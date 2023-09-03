// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAsiaHovdLocation  sync.Once
	cacheAsiaHovdLocation *time.Location
)

type AsiaHovd struct{}

func (AsiaHovd) Location() *time.Location {
	onceAsiaHovdLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Hovd")
		if err != nil {
			panic(err)
		}
		cacheAsiaHovdLocation = loc
	})
	return cacheAsiaHovdLocation
}
