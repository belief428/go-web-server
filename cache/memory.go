package logic

import (
	"encoding/json"
	"sync"
)

type Memory struct {
	cache       map[string]interface{}
	cacheList   []interface{}
	cacheSorted map[string][]*ScoreParams
	locker      *sync.RWMutex
}

func (this *Memory) Set(key string, value interface{}, expiration int) error {
	this.locker.Lock()
	defer this.locker.Unlock()
	this.cache[key] = value
	return nil
}

// Get
func (this *Memory) Get(key string) (string, error) {
	this.locker.RLock()
	defer this.locker.RUnlock()
	data, has := this.cache[key]

	if !has {
		return "", nil
	}
	_bytes, _ := json.Marshal(data)
	return string(_bytes), nil
}

// Del
func (this *Memory) Del(key string) error {
	this.locker.Lock()
	defer this.locker.Unlock()
	delete(this.cache, key)
	return nil
}
func (this *Memory) Run() error {
	return nil
}

func NewMemory() *Memory {
	return &Memory{
		cache:       make(map[string]interface{}, 0),
		cacheList:   make([]interface{}, 0),
		cacheSorted: make(map[string][]*ScoreParams, 0),
		locker:      new(sync.RWMutex),
	}
}
