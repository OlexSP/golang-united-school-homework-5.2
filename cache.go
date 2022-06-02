package cache

import "time"

type Cache struct {
	values map[string]item
}

type item struct {
	value    string
	deadLine *time.Time
}

// NewCache constructor func for Cache struct
func NewCache() Cache {
	return Cache{
		values: make(map[string]item),
	}
}

// Get methode  returns the value associated with the key  and the boolean ok
//(true if exists, false if not), if the deadline of the key/value pair has not been exceeded yet
func (c Cache) Get(key string) (string, bool) {
	t, ok := c.values[key] // comma ok time.Before/time.After
	if ok && (t.deadLine == nil || t.deadLine.After(time.Now())) {
		return t.value, true
	}
	return "", false
}

// Put places a value with an associated key into cache. Value put with this method never expired
//(have infinite deadline). Putting into the existing key should overwrite the value
func (c *Cache) Put(key, value string) { // needs mutex implementation!
	// interesting to check what is the item.deadLine parameter
	itemNew := item{value, nil}
	c.values[key] = itemNew
}

// Keys returns the slice of existing (non-expired keys)
func (c Cache) Keys() []string {
	var keySlice []string
	for k := range c.values {
		t, _ := c.values[k]
		valid := t.deadLine.After(time.Now())
		if t.deadLine == nil || valid {
			keySlice = append(keySlice, k)
		}
	}
	return keySlice
}

// PutTill does the same as Put, but the key/value pair should expire
//by given deadline
func (c *Cache) PutTill(key, value string, deadline time.Time) { // should implement mutex
	itemNew := item{value: value, deadLine: &deadline}
	c.values[key] = itemNew
}
