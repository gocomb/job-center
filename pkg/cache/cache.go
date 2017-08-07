package cache

import (
	"bytes"
	"errors"
	"sync"
	"time"
)

type Cache interface {
	// get cached value by key.
	Get(key string) interface{}
	// GetMulti is a batch version of Get.
	GetMulti(keys []string) []interface{}
	// set cached value with key and expire time.
	Put(key string, val interface{}) error
	// delete cached value by key.
	Delete(key string) error
	// increase cached int value by key, as a counter.
	Incr(key string) error
	// decrease cached int value by key, as a counter.
	Decr(key string) error
	// check if cached value exists or not.
	IsExist(key string) bool
	//Traverse the cache to get val with match string.
	MatchQueryKey(keyMatch string) []interface{}
	// clear all cache.
	ClearAll() error
}

// MemoryItem store memory cache item.
type MemoryItem struct {
	val         interface{}
	createdTime time.Time
}

// MemoryCache is Memory cache adapter.
// it contains a RW locker for safe map storage.
type MemoryCache struct {
	key interface{}
	sync.RWMutex
	dur   time.Duration
	items map[string]*MemoryItem
}

// NewMemoryCache returns a new MemoryCache.
func NewMemoryCache(key string) *MemoryCache {
	newCache := MemoryCache{key: key, items: make(map[string]*MemoryItem)}
	return &newCache
}

// Get cache from memory.
// if non-existed or expired, return nil.
func (bc *MemoryCache) Get(name string) interface{} {
	bc.RLock()
	defer bc.RUnlock()
	if itm, ok := bc.items[name]; ok {
		return itm.val
	}
	return nil
}

// GetMulti gets caches from memory.
// if non-existed or expired, return nil.
func (bc *MemoryCache) GetMulti(names []string) []interface{} {
	var rc []interface{}
	for _, name := range names {
		rc = append(rc, bc.Get(name))
	}
	return rc
}

// Put cache to memory.
// if lifespan is 0, it will be forever till restart.
func (bc *MemoryCache) Put(name string, value interface{}) error {
	bc.Lock()
	defer bc.Unlock()
	bc.items[name] = &MemoryItem{
		val:         value,
		createdTime: time.Now(),
	}
	return nil
}

// Delete cache in memory.
func (bc *MemoryCache) Delete(name string) error {
	bc.Lock()
	defer bc.Unlock()
	if _, ok := bc.items[name]; !ok {
		return errors.New("key not exist")
	}
	delete(bc.items, name)
	if _, ok := bc.items[name]; ok {
		return errors.New("delete key error")
	}
	return nil
}

// Incr increase cache counter in memory.
// it supports int,int32,int64,uint,uint32,uint64.
func (bc *MemoryCache) Incr(key string) error {
	bc.RLock()
	defer bc.RUnlock()
	itm, ok := bc.items[key]
	if !ok {
		return errors.New("key not exist")
	}
	switch itm.val.(type) {
	case int:
		itm.val = itm.val.(int) + 1
	case int32:
		itm.val = itm.val.(int32) + 1
	case int64:
		itm.val = itm.val.(int64) + 1
	case uint:
		itm.val = itm.val.(uint) + 1
	case uint32:
		itm.val = itm.val.(uint32) + 1
	case uint64:
		itm.val = itm.val.(uint64) + 1
	default:
		return errors.New("item val is not (u)int (u)int32 (u)int64")
	}
	return nil
}

// Decr decrease counter in memory.
func (bc *MemoryCache) Decr(key string) error {
	bc.RLock()
	defer bc.RUnlock()
	itm, ok := bc.items[key]
	if !ok {
		return errors.New("key not exist")
	}
	switch itm.val.(type) {
	case int:
		itm.val = itm.val.(int) - 1
	case int64:
		itm.val = itm.val.(int64) - 1
	case int32:
		itm.val = itm.val.(int32) - 1
	case uint:
		if itm.val.(uint) > 0 {
			itm.val = itm.val.(uint) - 1
		} else {
			return errors.New("item val is less than 0")
		}
	case uint32:
		if itm.val.(uint32) > 0 {
			itm.val = itm.val.(uint32) - 1
		} else {
			return errors.New("item val is less than 0")
		}
	case uint64:
		if itm.val.(uint64) > 0 {
			itm.val = itm.val.(uint64) - 1
		} else {
			return errors.New("item val is less than 0")
		}
	default:
		return errors.New("item val is not int int64 int32")
	}
	return nil
}

// IsExist check cache exist in memory.
func (bc *MemoryCache) IsExist(name string) bool {
	bc.RLock()
	defer bc.RUnlock()
	_, ok := bc.items[name]
	return ok
}

// ClearAll will delete all cache in memory.
func (bc *MemoryCache) ClearAll() error {
	bc.Lock()
	defer bc.Unlock()
	bc.items = make(map[string]*MemoryItem)
	return nil
}

// clearItems removes all the items which key in keys.
func (bc *MemoryCache) clearItems(keys []string) {
	bc.Lock()
	defer bc.Unlock()
	for _, key := range keys {
		delete(bc.items, key)
	}
}

//Traverse the cache to get val with match string.
func (bc *MemoryCache) MatchQueryKey(keyMatch string) []interface{} {
	var rc []interface{}
	for k := range bc.items {
		if bytes.Contains([]byte(k), []byte(keyMatch)) {
			rc = append(rc, bc.Get(k))
		}
	}
	return rc
}

//Register a cache, if return a ptr link interface
func Register(repoName string) Cache {
	return NewMemoryCache(repoName)
}
