// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAmericaMarigotLocation  sync.Once
	cacheAmericaMarigotLocation *time.Location
)

type AmericaMarigot struct{}

func (AmericaMarigot) Location() *time.Location {
	onceAmericaMarigotLocation.Do(func() {
		loc, err := time.LoadLocation("America/Marigot")
		if err != nil {
			panic(err)
		}
		cacheAmericaMarigotLocation = loc
	})
	return cacheAmericaMarigotLocation
}
