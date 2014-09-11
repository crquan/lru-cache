package lrucache

import (
	"container/list"
	"math"
	"sort"
)

type LRUCacheWithMap struct {
	hash  map[interface{}]*list.Element
	list  *list.List
	bound int
}

func NewLRUCacheWithMap() *LRUCacheWithMap {
	return &LRUCacheWithMap{
		hash: make(map[interface{}]*list.Element),
		list: list.New(),
		// by default is unbounded
		bound: math.MaxInt32,
	}
}

func (lruc *LRUCacheWithMap) SetBound(nb int) {
	if nb < 0 {
		return
	}

	for nb < lruc.list.Len() {
		e := lruc.list.Back()
		n := lruc.list.Remove(e).(*Node)
		delete(lruc.hash, n.key)
	}
	lruc.bound = nb
}

func (lruc *LRUCacheWithMap) Set(key, val interface{}) {
	ele, exists := lruc.hash[key]
	if exists {
		ele.Value.(*Node).val = val
	}

	if lruc.list.Len() == lruc.bound && lruc.bound > 0 {
		e := lruc.list.Back()
		n := e.Value.(*Node)
		delete(lruc.hash, n.key)
		lruc.list.Remove(e)
	}

	n := &Node{key, val}
	e := lruc.list.PushFront(n)
	lruc.hash[key] = e
}

func (lruc *LRUCacheWithMap) Get(key interface{}) (interface{}, bool) {
	ele, exists := lruc.hash[key]
	if !exists {
		return nil, false
	}
	lruc.list.MoveToFront(ele)
	return ele.Value.(*Node).val, true
}

func (lruc *LRUCacheWithMap) Peek(key interface{}) (interface{}, bool) {
	ele, exists := lruc.hash[key]
	if !exists {
		return nil, false
	}
	return ele.Value.(*Node).val, true
}

func (lruc *LRUCacheWithMap) Do(do func(key, value interface{})) {
	arr := make([]string, 0, len(lruc.hash))
	for key, _ := range lruc.hash {
		arr = append(arr, key.(string))
	}
	sort.Strings(arr)

	for _, key := range arr {
		e := lruc.hash[key]
		n := e.Value.(*Node)
		do(n.key, n.val)
	}
}
