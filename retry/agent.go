package retry

import (
	"github.com/appellative-ai/common/core"
	"github.com/appellative-ai/common/httpx"
	"github.com/appellative-ai/common/messaging"
	"github.com/appellative-ai/operatives/logger"
	"net/http"
	"sync/atomic"
	"time"
)

const (
	AgentName = "common:core:agent/operative/retry"
	duration  = time.Second * 30

	//collectiveName = "collective"
	host1Name = "host1"
	host2Name = "host2"
	routeName = "route"
)

type AgentT interface {
	messaging.Agent
	Exchange(req *http.Request) (resp *http.Response, err error)
}

type agentT struct {
	running atomic.Bool
	route   string
	host1   atomic.Value
	host2   atomic.Value
	timeout time.Duration

	exchange core.Exchange

	ticker   *messaging.Ticker
	emissary *messaging.Channel
}

func NewAgent(link map[string]string, timeout time.Duration) AgentT {
	a := new(agentT)
	a.running.Store(false)
	a.route = link[routeName]
	a.host1.Store(link[host1Name])
	a.host2.Store(link[host2Name])
	a.timeout = timeout

	a.exchange = httpx.Do

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()

	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return AgentName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case messaging.ConfigEvent:
		a.configure(m)
		return
	case messaging.StartupEvent:
		a.running.Store(true)
		return
	case messaging.ShutdownEvent:
		a.running.Store(false)
	}
}

// Exchange -
// TODO : failover if active host timeouts
func (a *agentT) Exchange(req *http.Request) (resp *http.Response, err error) {
	ctx, cancel := core.NewContext(req.Context(), a.timeout)
	defer cancel()

	req = req.Clone(ctx)
	err = newURL(req, a.host1.Load().(string))
	if err != nil {
		return httpx.NewResponse(http.StatusBadRequest, nil, nil), err
	}
	start := time.Now().UTC()
	resp, err = a.exchange(req)
	logger.Agent.LogEgress(start, time.Since(start), a.route, req, resp, a.timeout)
	if err != nil {
		return
	}
	err = httpx.TransformBody(resp)
	return
}
