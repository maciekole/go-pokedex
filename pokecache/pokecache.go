package pokecache

import "time"

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Interval int64
	Data     map[string]CacheEntry
}

func (cache *Cache) NewCache(interval int64) (*Cache, error) {
	newCache := &Cache{
		Interval: interval,
	}
	err := newCache.reapLoop()
	if err != nil {
		return nil, err
	}
	return newCache, nil
}

func (cache *Cache) Add(key string, val []byte) error {
	return nil
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	return nil, false
}

func (cache *Cache) reapLoop() error {
	return nil
}
