package mansbridge

import (
	"sync"
)

//based on http://play.golang.org/p/_FvECoFvhq

type peerList struct {
	sync.RWMutex
	set map[string]bool
}

func newPeerList() *peerList {
	return &peerList{set: make(map[string]bool)}
}

func (set *peerList) add(i string) bool {
	set.Lock()
	_, found := set.set[i]
	set.set[i] = true
	set.Unlock()
	return !found
}

func (set *peerList) get(i string) bool {
	set.RLock()
	_, found := set.set[i]
	set.Unlock()
	return found
}

func (set *peerList) getAll() []string {
	set.RLock()
	keys := make([]string, 0, len(set.set))
	for k := range set.set {
		keys = append(keys, k)
	}
	set.RUnlock()
	return keys
}

func (set *peerList) remove(i string) {
	set.Lock()
	delete(set.set, i)
	set.Unlock()
}
