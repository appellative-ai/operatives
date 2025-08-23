package retry

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
		s := l[primaryHost]
		if s != "" {
			a.primary.Store(s)
		}
		s = l[secondaryHost]
		if s != "" {
			a.secondary.Store(s)
		}
		return
	}
}
