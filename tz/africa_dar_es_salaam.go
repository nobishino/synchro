// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAfricaDar_es_SalaamLocation  sync.Once
	cacheAfricaDar_es_SalaamLocation *time.Location
)

type AfricaDar_es_Salaam struct{}

func (AfricaDar_es_Salaam) Location() *time.Location {
	onceAfricaDar_es_SalaamLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Dar_es_Salaam")
		if err != nil {
			panic(err)
		}
		cacheAfricaDar_es_SalaamLocation = loc
	})
	return cacheAfricaDar_es_SalaamLocation
}
