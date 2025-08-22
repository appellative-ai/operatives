package concurrency

import (
	"github.com/appellative-ai/common/core"
	"github.com/appellative-ai/common/messaging"
)

func (a *agentT) configure(m *messaging.Message) {
	if m == nil || m.Name != messaging.ConfigEvent {
		return
	}
	if ex, ok := messaging.ConfigContent[core.Exchange](m); ok && ex != nil {
		if !a.running.Load() {
			a.exchange = ex
		}
		return
	}
	if l, ok := messaging.ConfigContent[map[string]string](m); ok && len(l) > 0 {
		s := l[host1Name]
		if s != "" {
			a.host1.Store(s)
		}
		s = l[host2Name]
		if s != "" {
			a.host2.Store(s)
		}
		return
	}
}
