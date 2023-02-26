package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value any) bool
	Get(key Key) (any, bool)
	Clear() Cache
	Len() int
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	lock     sync.Mutex
}

func (s *lruCache) Set(key Key, value any) (inCache bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, inCache = s.items[key]; inCache {
		s.items[key].Value = value
		if s.items[key] != s.queue.Front() {
			s.queue.MoveToFront(s.items[key].Prev.Next)
		}
	} else {
		if s.queue.Len() >= s.capacity {
			k := s.queue.Back().Key
			s.queue.Remove(s.queue.Back())
			delete(s.items, k)
		}
		s.items[key] = s.queue.PushFront(value)
		s.items[key].Key = key
	}
	return inCache
}

func (s *lruCache) Get(key Key) (value any, inCache bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, inCache = s.items[key]; inCache {
		value = s.items[key].Value
		if s.items[key] != s.queue.Front() {
			s.queue.MoveToFront(s.items[key].Prev.Next)
		}
	}
	return value, inCache
}

func (s *lruCache) Clear() Cache {
	return NewCache(s.capacity)
}

func (s *lruCache) Len() int {
	return s.queue.Len()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
