package main

import (
	"container/list"
	"fmt"
)

type LRUCallback func(key, val string)

type KV struct {
	key string
	val string
}

type LRUCache struct {
	size   int
	store  map[string]*list.Element
	elList *list.List
	cbList []LRUCallback
}

func (c *LRUCache) RegisterCallback(cb LRUCallback) {
	c.cbList = append(c.cbList, cb)
}

func (c *LRUCache) Add(key, val string) {
	e, ok := c.store[key]
	if !ok {
		e = c.elList.PushFront(&KV{key: key, val: val})
		c.store[key] = e
	} else {
		kv := e.Value.(*KV)
		kv.val = val
		c.elList.MoveToFront(e)
	}

	if len(c.store) <= c.size {
		return
	}

	last := c.elList.Back()
	c.elList.Remove(last)
	kv := last.Value.(*KV)
	delete(c.store, kv.key)
	for _, cb := range c.cbList {
		cb(kv.key, kv.val)
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	e, ok := c.store[key]
	val := ""
	if ok {
		kv := e.Value.(*KV)
		c.elList.MoveToFront(e)
		val = kv.val
	}

	return val, ok
}

func main() {
	l := list.New()
	fmt.Printf("list: %+v", l)
}
