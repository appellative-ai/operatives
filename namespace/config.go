package namespace

import (
	"github.com/appellative-ai/common/messaging"
)

func (a *agentT) configure(m *messaging.Message) {
	if m == nil || m.Name != messaging.ConfigEvent {
		return
	}
	if l, ok := messaging.ConfigContent[[]map[string]string](m); ok && len(l) > 0 {
		var links []map[string]string
		links = append(links, l...)
		a.links.Store(&links)
		return
	}
}
