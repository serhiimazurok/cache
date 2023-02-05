package cache

type Cache struct {
	kv map[string]any
}

func New() *Cache {
	return &Cache{
		kv: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	c.kv[key] = value
}

func (c *Cache) Get(key string) any {
	return c.kv[key]
}

func (c *Cache) Delete(key string) {
	delete(c.kv, key)
}
