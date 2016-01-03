package mansbridge

type NewsItem struct {
	AgentId string
	Item    interface{}
}

//buffer specifys the number of news items we
//will hold in between getNews reads
func NewAgent(agentId string, buffer int) *Agent {
	r := make(chan NewsItem)
	l := make(chan NewsItem, buffer)
	return &Agent{id: agentId, RemoteNews: r, localNews: l}
}

type Agent struct {
	localNews  chan NewsItem
	RemoteNews chan NewsItem
	id         string
}

func (a *Agent) newsUpdate(items []NewsItem) {
	for _, item := range items {
		a.RemoteNews <- item
	}
}

func (a *Agent) AddNews(item interface{}) {
	a.localNews <- NewsItem{AgentId: a.id, Item: item}
}

func (a *Agent) getNews() NewsItem {
	select {
	case news := <-a.localNews:
		return news
	default:
		return NewsItem{}

	}
}

func (a *Agent) Stop() {
	close(a.localNews)
	close(a.RemoteNews)
}
