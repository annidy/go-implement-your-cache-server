package cache

import "log"

func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache()
	}
	if typ == "inmemory2" {
		c = newInMemoryCache2()
	}
	if typ == "rocksdb" {
		c = newRocksdbCache()
	}
	if c == nil {
		panic("unknown cache type " + typ)
	}
	log.Println(typ, "ready to serve")
	return c
}
