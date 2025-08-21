package namespace

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/appellative-ai/center/exchange"
	"github.com/appellative-ai/center/template"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/postgres/request"
	"github.com/appellative-ai/postgres/retrieval"
	"net/http"
	"time"
)

const (
	AgentName = "common:core:agent/namespace/center"
	duration  = time.Second * 30
	timeout   = time.Second * 4

	retrievalPath    = "/namespace/retrieval"
	relationPath     = "/namespace/relation"
	requestThingPath = "/namespace/request/thing"
	requestLinkPath  = "/namespace/request/link"
)

type agentT struct {
	running bool
	timeout time.Duration

	ticker   *messaging.Ticker
	emissary *messaging.Channel

	retriever *retrieval.Interface
	requester *request.Interface
	processor template.Agent
}

// init - register an agent constructor
func init() {
	exchange.RegisterConstructor(AgentName, func() messaging.Agent {
		return newAgent()
	})
}

func newAgent() *agentT {
	a := new(agentT)
	a.timeout = timeout
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()

	a.retriever = retrieval.Retriever
	a.requester = request.Requester
	a.processor = template.NewAgent(retrieval.Retriever)
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
	go emissaryAttend(a)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

// Link - chainable exchange
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		ctx, cancel := httpx.NewContext(nil, a.timeout)
		defer cancel()
		var buf *bytes.Buffer

		switch req.URL.Path {
		case retrievalPath:
			buf, err = retrievalRequest(ctx, a.retriever, a.processor, req)
		case relationPath:
			buf, err = relationRequest(ctx, a.retriever, a.processor, req)
		case requestThingPath:
			_, err = thingRequest(ctx, a.requester, req.URL.Query())
		case requestLinkPath:
			_, err = linkRequest(ctx, a.requester, req.URL.Query())
		default:
			return httpx.NewResponse(http.StatusBadRequest, nil, nil), errors.New(fmt.Sprintf("path is invalid [%v]", req.URL.Path))
		}
		if err != nil {
			return httpx.NewResponse(http.StatusInternalServerError, nil, nil), err
		}
		return httpx.NewResponse(http.StatusOK, nil, buf), nil
	}
}
