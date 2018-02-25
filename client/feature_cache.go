package client

import (
	"bytes"
	"encoding/gob"
	"sync"
)

// FeatureCache is a in-memory threadsafe cache for Features.
type FeatureCache struct {
	features map[string]bool
	lock     sync.RWMutex
}

// NewFeatureCache creates a new FeatureCache.
func NewFeatureCache() *FeatureCache {
	return &FeatureCache{
		features: make(map[string]bool),
	}
}

// Add adds a feature to the cache.
func (fc *FeatureCache) Add(feature string, status bool) {
	fc.lock.Lock()
	defer fc.lock.Unlock()

	fc.features[feature] = status
}

// AddAll adds a list of features to the cache.
func (fc *FeatureCache) AddAll(features map[string]bool) {
	fc.lock.Lock()
	defer fc.lock.Unlock()

	for _, feature := range features {
		fc.features[feature] = true
	}
}

// Get gets a Feature from the cache if it exits.
func (fc *FeatureCache) Get(name string) bool {
	fc.lock.RLock()
	defer fc.lock.RUnlock()

	feature, ok := fc.features[name]
	return ok && feature
}

func (fc *FeatureCache) GetAll() map[string]bool {
	fc.lock.RLock()
	defer fc.lock.RUnlock()

	// deep clone features map
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	err := enc.Encode(fc.features)
	if err != nil {
		panic(err)
	}

	var deepCopy map[string]bool
	err = dec.Decode(&deepCopy)
	if err != nil {
		panic(err)
	}

	return deepCopy
}
