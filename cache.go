package cache

import "time"

type Cache struct {
	values map[string]Item
}

type Item struct {
	value    string
	deadLine *time.Time
}

// NewCache constructor func for Cache struct
func NewCache() Cache {
	return Cache{
		values: map[string]Item{},
	}
}

// Get methode  returns the value associated with the key  and the boolean ok
//(true if exists, false if not), if the deadline of the key/value pair has not been exceeded yet
func (c *Cache) Get(key string) (string, bool) {
	t, ok := c.values[key] // comma ok time.Before/time.After

	if ok && (t.deadLine == nil || t.deadLine.After(time.Now())) {
		return t.value, true
	}
	return "", false
}

// Put places a value with an associated key into cache. Value put with this method never expired
//(have infinite deadline). Putting into the existing key should overwrite the value
func (c *Cache) Put(key, value string) {
	c.values[key] = Item{value, nil}
}

// Keys returns the slice of existing (non-expired keys)
func (c *Cache) Keys() []string {
	var keySlice []string
	for k, v := range c.values {
		if v.deadLine == nil || v.deadLine.After(time.Now()) {
			keySlice = append(keySlice, k)
		}
	}
	return keySlice
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.values[key] = Item{value, &deadline}
}
