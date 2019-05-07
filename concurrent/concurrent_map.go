package concurrent

import (
	"sync"
)

type ForEachHandler func(key string, value interface{})

type AppendHandler func(cm MapType) MapType

type MapType map[string]interface{}

type ConcurrentMap struct {
	m    MapType
	lock sync.Mutex
}

// NewConcurrentMap It creates a new concurrent map obeject
func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(MapType),
	}
}

// Add It adds a new object in map
func (cm *ConcurrentMap) Add(key string, value interface{}) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	cm.m[key] = value
}

// Get It returns an object if the object is exist
func (cm *ConcurrentMap) Get(key string, defaultValue interface{}) interface{} {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	if value, ok := cm.m[key]; ok {
		return value
	}

	return defaultValue
}

// HasKey It checks whether the map has given key or not
func (cm *ConcurrentMap) HasKey(key string) bool {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	_, ok := cm.m[key]
	return ok
}

// Append It appends the value into existing object like as append a value into array
func (cm *ConcurrentMap) Append(handler AppendHandler) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	cm.m = handler(cm.m)
}

// ForEach It iterates all objects from map
func (cm *ConcurrentMap) ForEach(handler ForEachHandler) {
	for key, value := range cm.m {
		handler(key, value)
	}
}

// Interface It returns whole map in interface type
func (cm *ConcurrentMap) Interface() interface{} {
	return cm.m
}

// Length It returns the map size
func (cm *ConcurrentMap) Length() int {
	return len(cm.m)
}
