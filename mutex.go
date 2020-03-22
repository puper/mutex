package mutex

import (
	"sync"
)

var (
	defaultMutexManager = NewMutexManager()
)

type Mutex struct {
	sync.RWMutex
	locks int64
}

func NewMutexManager() *MutexManager {
	return &MutexManager{
		mutexes: map[string]*Mutex{},
	}
}

type MutexManager struct {
	mutex   sync.Mutex
	mutexes map[string]*Mutex
}

func Lock(key string) {
	defaultMutexManager.Lock(key)
}
func (this *MutexManager) Lock(key string) {
	this.mutex.Lock()
	if _, ok := this.mutexes[key]; !ok {
		this.mutexes[key] = &Mutex{}
	}
	this.mutexes[key].locks++
	this.mutex.Unlock()
	this.mutexes[key].Lock()

}
func Unlock(key string) {
	defaultMutexManager.Unlock(key)
}
func (this *MutexManager) Unlock(key string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if _, ok := this.mutexes[key]; ok {
		this.mutexes[key].Unlock()
		this.mutexes[key].locks--
		if this.mutexes[key].locks == 0 {
			delete(this.mutexes, key)
		}
	} else {
		panic("unlock of unlocked mutex")
	}
}

func RLock(key string) {
	defaultMutexManager.RLock(key)
}
func (this *MutexManager) RLock(key string) {
	this.mutex.Lock()
	if _, ok := this.mutexes[key]; !ok {
		this.mutexes[key] = &Mutex{}
	}
	this.mutexes[key].locks++
	this.mutex.Unlock()
	this.mutexes[key].RLock()

}

func RUnlock(key string) {
	defaultMutexManager.RUnlock(key)
}
func (this *MutexManager) RUnlock(key string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if _, ok := this.mutexes[key]; ok {
		this.mutexes[key].RUnlock()
		this.mutexes[key].locks--
		if this.mutexes[key].locks == 0 {
			delete(this.mutexes, key)
		}
	} else {
		panic("unlock of unlocked mutex")
	}
}
