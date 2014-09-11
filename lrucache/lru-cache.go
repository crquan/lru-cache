package lrucache

import (
	"./omap"
	"container/list"
	"math"
)

type Node struct {
	key, val interface{}
}

type LRUCache struct {
	hash  *omap.Map
	list  *list.List
	bound int
}

func New() *LRUCache {
	return &LRUCache{
		hash: omap.NewStringOrderedMap(),
		list: list.New(),
		// by default is unbounded
		bound: math.MaxInt32,
	}
}

func (lruc *LRUCache) SetBound(nb int) {
	if nb < 0 {
		return
	}

	for nb < lruc.list.Len() {
		e := lruc.list.Back()
		n := lruc.list.Remove(e).(*Node)
		lruc.hash.Delete(n.key)
	}
	lruc.bound = nb
}

func (lruc *LRUCache) Set(key, val interface{}) {
	ele, exists := lruc.hash.Find(key)
	if exists {
		ele.(*list.Element).Value.(*Node).val = val
	}

	if lruc.list.Len() == lruc.bound && lruc.bound > 0 {
		e := lruc.list.Back()
		n := e.Value.(*Node)
		lruc.hash.Delete(n.key)
		lruc.list.Remove(e)
	}

	n := &Node{key, val}
	e := lruc.list.PushFront(n)
	lruc.hash.Insert(key, e)
}

func (lruc *LRUCache) Get(key interface{}) (interface{}, bool) {
	ele, found := lruc.hash.Find(key)
	if !found {
		return nil, false
	}
	lruc.list.MoveToFront(ele.(*list.Element))
	return ele.(*list.Element).Value.(*Node).val, true
}

func (lruc *LRUCache) Peek(key interface{}) (interface{}, bool) {
	ele, exists := lruc.hash.Find(key)
	if !exists {
		return nil, false
	}
	return ele.(*list.Element).Value.(*Node).val, true
}

func (lruc *LRUCache) Do(do func(key, value interface{})) {
	lruc.hash.Do(func(key, value interface{}) {
		do(key, value.(*list.Element).Value.(*Node).val)
	})
}
