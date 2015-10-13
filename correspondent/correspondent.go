package correspondent

import (
	"github.com/nstehr/go-mansbridge/agent"
	"github.com/nstehr/go-mansbridge/collections"
	"github.com/nstehr/go-mansbridge/wire"
	"log"
	"math/rand"
	"time"
)

const (
	cacheSize = 10
)

type Correspondent struct {
	agent              agent.Agent
	cache              *collections.Cache
	peers              *collections.PeerList
	expectingReplyFrom string
	wireService        wire.WireService
	done               chan bool
	refreshInterval    time.Duration
}

func NewCorrespondent(a agent.Agent, seedIp string, wireService wire.WireService, refreshInterval int) *Correspondent {
	peers := collections.NewPeerList()
	peers.Add(seedIp)
	doneChan := make(chan bool)
	c := Correspondent{agent: a,
		cache:              collections.NewCache(cacheSize),
		wireService:        wireService,
		expectingReplyFrom: "",
		peers:              peers,
		done:               doneChan,
		refreshInterval:    time.Duration(refreshInterval)}

	go c.listenForRemoteUpdates()

	return &c
}

func (c *Correspondent) listenForRemoteUpdates() {
	for msg := range c.wireService.GetNews() {
		//if we aren't expecting a reply, we were
		//the randomly selected remote, so send our
		//cache over
		if c.expectingReplyFrom != msg.Source {
			log.Println("Msg from: " + msg.Source)
			go c.wireService.SendNews(msg.Source, c.cache.GetEntries())
		} else {
			//this is a reply, so we can clear that we got it
			log.Println("Reply from: " + msg.Source)
			c.expectingReplyFrom = ""
		}
		var remoteNews []agent.NewsItem
		for _, entry := range msg.Entries {
			if entry.IpAddress != c.wireService.GetAddress() {
				//collect the non-local news to pass to the agent
				remoteNews = append(remoteNews, entry.News)
				//update the list of peers that we can send to
				c.peers.Add(entry.IpAddress)
			}
		}
		//pass the entries up to the agent
		c.agent.NewsUpdate(remoteNews)
		//add and refresh cache
		c.cache.AddEntries(msg.Entries...)
		c.cache.Resize()
	}
}

func (c *Correspondent) StartReporting() {
	tick := time.NewTicker(time.Second * c.refreshInterval).C

	for {
		select {
		case <-tick:
			log.Println("Checking for new news...")
			//step 1, add news to cache
			entry := collections.Entry{IpAddress: c.wireService.GetAddress(),
				Timestamp: time.Now(),
				News:      c.agent.GetNews()}
			c.cache.AddEntries(entry)
			//step 2, find a random peer
			peer := findPeer(c.peers.GetAll())
			//step 3, send cache to peer
			log.Println("Sending cache to: " + peer)
			c.wireService.SendNews(peer, c.cache.GetEntries())
			//keep track of who we sent to, so we can expect a response
			c.expectingReplyFrom = peer

		case <-c.done:
			log.Println("Done")
			return
		}
	}
}

func (c *Correspondent) StopReporting() {
	c.done <- true
}

func findPeer(peers []string) string {
	rand.Seed(time.Now().Unix())
	idx := rand.Intn(len(peers)-0) + 0
	return peers[idx]
}
