package main

import (
	"fmt"
	"sync"
)

func MakeSyncMap(capacity int) SyncMap {
	return SyncMap{
		dict: make(map[string]interface{}, capacity),
	}
}

// Реализовать конкурентную запись данных в map.
type SyncMap struct {
	dict map[string]interface{}
	mu   sync.RWMutex
}

func (sm *SyncMap) Set(key string, val interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.dict[key] = val
}

func (sm *SyncMap) Get(key string) (interface{}, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.dict[key]
	return val, ok
}

func (sm *SyncMap) Del(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.dict, key)
}

// P.S.: Вместо этого можно использовать sync.Map

func main() {
	smap := MakeSyncMap(1000)

	wg := sync.WaitGroup{}
	for i := 0; i < 1_000; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			smap.Set(fmt.Sprint(index), fmt.Sprint(index))
		}(i)
	}

	wg.Wait()

	fmt.Println(smap.dict)
}
