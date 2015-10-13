package agent

type NewsItem struct {
	AgentId string
	Item    interface{}
}

type Agent interface {
	GetNews() NewsItem
	NewsUpdate(item []NewsItem)
}
