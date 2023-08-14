package cache

import "sync"

type inMemoryCache2 struct {
	c sync.Map
	Stat
}

func (c *inMemoryCache2) Set(k string, v []byte) error {
	c.c.Store(k, v)
	c.add(k, v)
	return nil
}

func (c *inMemoryCache2) Get(k string) ([]byte, error) {
	v, ok := c.c.Load(k)
	if !ok {
		return nil, nil
	}
	return []byte(v.(string)), nil
}

func (c *inMemoryCache2) Del(k string) error {
	c.c.Delete(k)
	return nil
}

func (c *inMemoryCache2) GetStat() Stat {
	return c.Stat
}

func newInMemoryCache2() *inMemoryCache2 {
	return &inMemoryCache2{sync.Map{}, Stat{}}
}
