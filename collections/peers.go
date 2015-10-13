package collections

import (
	"sync"
)

//based on http://play.golang.org/p/_FvECoFvhq

type PeerList struct {
	sync.RWMutex
	set map[string]bool
}

func NewPeerList() *PeerList {
	return &PeerList{set: make(map[string]bool)}
}

func (set *PeerList) Add(i string) bool {
	set.Lock()
	_, found := set.set[i]
	set.set[i] = true
	set.Unlock()
	return !found
}

func (set *PeerList) Get(i string) bool {
	set.RLock()
	_, found := set.set[i]
	set.Unlock()
	return found
}

func (set *PeerList) GetAll() []string {
	set.RLock()
	keys := make([]string, 0, len(set.set))
	for k := range set.set {
		keys = append(keys, k)
	}
	set.RUnlock()
	return keys
}

func (set *PeerList) Remove(i string) {
	set.Lock()
	delete(set.set, i)
	set.Unlock()
}
