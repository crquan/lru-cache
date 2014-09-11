package lrucache

import (
	_ "container/list"
	"fmt"
	"testing"
)

func printData(get func(interface{}) (interface{}, bool), key interface{}) {
	if val, exists := get(key); !exists {
		fmt.Println("NULL")
	} else {
		fmt.Println(val)
	}
}

func ExampleLRUCache() {
	lru := NewLRUCacheWithMap()
	lru.SetBound(2)
	lru.Set("a", 2)
	lru.Set("b", 4)
	printData(lru.Get, "b")
	printData(lru.Peek, "a")
	lru.Set("c", 5)
	printData(lru.Get, "a")
	lru.Do(func(key, val interface{}) {
		fmt.Println(key, val)
	})
	// Output:
	// 4
	// 2
	// NULL
	// b 4
	// c 5
}

func BenchmarkLRUCache(b *testing.B) {
	lru := New()
	// fmt.Println("testing cycle:", b.N)

	for i := 1; i <= b.N; i++ {
		lru.Set(fmt.Sprint(i), i)
	}
	for i := 200; i >= 100; i-- {
		lru.Get(fmt.Sprint(i))
	}
	lru.SetBound(100)
	for i := 100; i >= 1; i-- {
		lru.Set(fmt.Sprint(i), i)
	}
	for i := 1; i <= b.N; i++ {
		lru.Get(fmt.Sprint(i))
	}
	for i := 1; i <= b.N; i++ {
		lru.Peek(fmt.Sprint(i))
	}
	// fmt.Println(lru.bound, lru.list.Len())

	sum0 := 0
	csum := func(key, val interface{}) {
		sum0 += len(key.(string)) + val.(int)
	}
	lru.Do(csum)
	// fmt.Println(b.N, sum0)
}

func BenchmarkLRUCacheWithMap(b *testing.B) {
	lru := NewLRUCacheWithMap()
	// fmt.Println("testing cycle:", b.N)

	for i := 1; i <= b.N; i++ {
		lru.Set(fmt.Sprint(i), i)
	}
	for i := 200; i >= 100; i-- {
		lru.Get(fmt.Sprint(i))
	}
	lru.SetBound(100)
	for i := 100; i >= 1; i-- {
		lru.Set(fmt.Sprint(i), i)
	}
	for i := 1; i <= b.N; i++ {
		lru.Get(fmt.Sprint(i))
	}
	for i := 1; i <= b.N; i++ {
		lru.Peek(fmt.Sprint(i))
	}
	// fmt.Println(lru.bound, lru.list.Len())

	sum0 := 0
	csum := func(key, val interface{}) {
		sum0 += len(key.(string)) + val.(int)
	}
	lru.Do(csum)
	// fmt.Println(b.N, sum0)
}
