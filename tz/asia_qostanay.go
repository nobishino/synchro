// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAsiaQostanayLocation  sync.Once
	cacheAsiaQostanayLocation *time.Location
)

type AsiaQostanay struct{}

func (AsiaQostanay) Location() *time.Location {
	onceAsiaQostanayLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Qostanay")
		if err != nil {
			panic(err)
		}
		cacheAsiaQostanayLocation = loc
	})
	return cacheAsiaQostanayLocation
}
