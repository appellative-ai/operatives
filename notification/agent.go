package notification

import (
	"fmt"
	"github.com/appellative-ai/common/core"
	"github.com/appellative-ai/common/messaging"
	"github.com/appellative-ai/postgres/request"
	"github.com/appellative-ai/postgres/retrieval"
	"net/http"
	"time"
)

const (
	AgentName = "common:core:agent/notification/center"
	duration  = time.Second * 30
	timeout   = time.Second * 4
)

type agentT struct {
	running bool
	timeout time.Duration

	ticker   *messaging.Ticker
	emissary *messaging.Channel

	retriever *retrieval.Interface
	requester *request.Interface
}

func newAgent() *agentT {
	a := new(agentT)
	a.timeout = timeout
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.retriever = retrieval.Retriever
	a.requester = request.Requester

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
		if a.running {
			return
		}
		messaging.UpdateContent[time.Duration](m, &a.timeout)
		return
	case messaging.StartupEvent:
		if a.running {
			return
		}
		a.running = true
		a.run()
		return
	case messaging.ShutdownEvent:
		if !a.running {
			return
		}
		a.running = false
	}
	switch m.Channel() {
	case messaging.ChannelControl, messaging.ChannelEmissary:

		a.emissary.C <- m
	default:
		fmt.Printf("limiter - invalid channel %v\n", m)
	}
}

// Run - run the agent
func (a *agentT) run() {
	//go emissaryAttend(a)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

// Link - chainable exchange
func (a *agentT) Link(next core.Exchange) core.Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		//ctx, cancel := httpx.NewContext(nil, a.timeout)
		//defer cancel()

		/*
			buf, err1 := a.retriever.Marshal(ctx, ""thingQueryName, "select * from thing", nil)
			if err1 != nil {
				return httpx.NewResponse(messaging.StatusExecError, nil, err1), err1
			}
			return httpx.NewResponse(http.StatusOK, nil, buf), nil

		*/
		return
	}
}
