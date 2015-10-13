package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"github.com/nstehr/go-mansbridge/agent"
	"github.com/nstehr/go-mansbridge/correspondent"
	"github.com/nstehr/go-mansbridge/wire"
	"log"
)

type TestAgent struct {
	id string
}

func (t TestAgent) GetNews() agent.NewsItem {
	return agent.NewsItem{Item: "asdfasfdasfd", AgentId: t.id}
}
func (t TestAgent) NewsUpdate(item []agent.NewsItem) {

}

func main() {
	port := flag.Int("port", 10001, "Port To Listen For News On")
	seed := flag.String("seed", "localhost", "Initial Known Peer")
	flag.Parse()

	agentId := generateId()
	log.Println("Agent: " + agentId)

	encoder := wire.GobWireEncoder{}
	wireService := wire.NewUdpWireService(*port, encoder)

	c := correspondent.NewCorrespondent(TestAgent{id: agentId}, *seed, wireService)

	c.StartReporting()

}

func generateId() string {
	size := 32

	rb := make([]byte, size)
	_, err := rand.Read(rb)

	if err != nil {
		log.Println(err)
	}

	rs := base64.URLEncoding.EncodeToString(rb)

	return rs
}
