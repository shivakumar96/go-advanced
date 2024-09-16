package customcontainer

import (
	"container/list"
	"errors"
	"fmt"
)

/**
*  Implenets LRU cache
 */

var (
	ErrKEY_NOT_FOUND = errors.New("KEY NOT FOUND")
)

type LRUType interface {
	~int | ~string
}

type pair[K LRUType, V LRUType] struct {
	key   K
	value V
}

type LRUmap[K LRUType] map[K]*list.Element

// using struct embading
type LRUCache[K LRUType, V LRUType] struct {
	LRUmap[K]
	dlist    *list.List
	capacity int
}

func NewLRUCache[K LRUType, V LRUType](cpapcity int) *LRUCache[K, V] {
	cache := &LRUCache[K, V]{
		capacity: cpapcity,
		dlist:    list.New(),
		LRUmap:   make(LRUmap[K], cpapcity),
	}

	return cache
}

func (cache *LRUCache[K, V]) Put(key K, val V) {
	ele, found := cache.LRUmap[key]
	if found {
		casted_ele := ele.Value.(*pair[K, V])
		casted_ele.value = val
		cache.dlist.MoveToFront(ele)
		//fmt.Println(casted_ele)
		//fmt.Println(ele)
	} else {
		if cache.dlist.Len() == cache.capacity {
			cache.dlist.Remove(cache.dlist.Back())
		}
		ele = cache.dlist.PushFront(&pair[K, V]{key: key, value: val})
		cache.LRUmap[key] = ele
	}
}

func (cache *LRUCache[K, V]) Get(key K) (V, error) {
	ele, found := cache.LRUmap[key]
	if !found {
		return *new(V), ErrKEY_NOT_FOUND
	}
	ans := ele.Value.(*pair[K, V])
	cache.dlist.MoveToFront(ele)
	return ans.value, nil
}

func (cache *LRUCache[K, V]) PrintList() {
	lis := cache.dlist
	fmt.Println(cache.LRUmap)
	for ele := lis.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value.(*pair[K, V]))
	}
}
