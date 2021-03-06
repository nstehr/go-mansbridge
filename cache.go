package mansbridge

import (
	"sort"
	"sync"
	"time"
)

type cache struct {
	sync.RWMutex
	entries  Entries
	capacity int
}

type Entry struct {
	IpAddress string
	Timestamp time.Time
	News      NewsItem
}

type Entries []Entry

func (entries Entries) Len() int {
	return len(entries)
}

func (entries Entries) Less(i, j int) bool {
	return entries[i].Timestamp.After(entries[j].Timestamp)
}

func (entries Entries) Swap(i, j int) {
	entries[i], entries[j] = entries[j], entries[i]
}

func newCache(capacity int) *cache {
	var entries Entries
	return &cache{entries: entries, capacity: capacity}
}

func (c *cache) addEntries(newEntries ...Entry) {
	c.Lock()
	c.entries = append(c.entries, newEntries...)
	c.Unlock()
}

func (c *cache) resize() {
	c.Lock()
	//definitely a more efficient way to hanlding this.
	if len(c.entries) > c.capacity {
		sort.Sort(c.entries)
		c.entries = c.entries[0:c.capacity]
	}
	c.Unlock()
}

func (c *cache) getEntries() []Entry {
	c.Lock()
	copyEntries := make([]Entry, len(c.entries))
	copy(copyEntries, c.entries)
	c.Unlock()
	return copyEntries

}
