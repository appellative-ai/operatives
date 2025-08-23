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

	primaryHostName   = "primary"
	secondaryHostName = "secondary"
	routeName         = "X-Route-Name"
)

type AgentT interface {
	messaging.Agent
	Exchange(req *http.Request) (resp *http.Response, err error)
}

type agentT struct {
	running   atomic.Bool
	primary   atomic.Value
	secondary atomic.Value
	timeout   time.Duration

	exchange core.Exchange

	ticker   *messaging.Ticker
	emissary *messaging.Channel
}

func NewAgent(timeout time.Duration) AgentT {
	return newAgent(timeout)
}

func newAgent(timeout time.Duration) *agentT {
	a := new(agentT)
	a.running.Store(false)
	a.primary.Store("")
	a.secondary.Store("")
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
func (a *agentT) Exchange(req *http.Request) (resp *http.Response, err error) {
	route := req.Header.Get(routeName)
	req.Header.Del(routeName)
	resp, err = a.do(req, a.primary.Load().(string), route)
	if err != nil || (resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusGatewayTimeout) {
		resp, err = a.do(req, a.secondary.Load().(string), route)
	}

	return
}

func (a *agentT) do(req *http.Request, host, route string) (resp *http.Response, err error) {
	ctx, cancel := core.NewContext(req.Context(), a.timeout)
	defer cancel()

	req = req.Clone(ctx)
	err = newURL(req, host)
	if err != nil {
		return httpx.NewResponse(http.StatusBadRequest, nil, nil), err
	}
	start := time.Now().UTC()
	resp, err = a.exchange(req)
	if err != nil {
		logger.Agent.LogEgress(start, time.Since(start), route, req, resp, a.timeout)
		return
	}
	err = httpx.TransformBody(resp)
	logger.Agent.LogEgress(start, time.Since(start), route, req, resp, a.timeout)
	return
}
