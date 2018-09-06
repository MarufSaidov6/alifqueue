package types

import (
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex                      //Security Access
	DefaultExpiration time.Duration   //Cache LifeTime
	CleanupInterval   time.Duration   //Cache CleanUp
	Items             map[string]Item //Cache Elements,*string->ID
}

type Item struct {
	Value      interface{} //Instances
	Created    time.Time
	Expiration int64
}
