package namespace

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/appellative-ai/common/core"
	"github.com/appellative-ai/common/httpx"
	"github.com/appellative-ai/common/messaging"
	"github.com/appellative-ai/operatives/retry"
	"github.com/appellative-ai/operatives/template"
	"github.com/appellative-ai/postgres/request"
	"github.com/appellative-ai/postgres/retrieval"
	"net/http"
	"sync/atomic"
	"time"
)

const (
	AgentName = "common:core:agent/namespace/center"
	duration  = time.Second * 30
	timeout   = time.Second * 4

	collectiveName    = "collective"
	primaryHostName   = "primary"
	secondaryHostName = "secondary"
	routeName         = "route"

	retrievalPath    = "/namespace/retrieval"
	relationPath     = "/namespace/relation"
	requestThingPath = "/namespace/request/thing"
	requestLinkPath  = "/namespace/request/link"
)

var (
	Agent messaging.Agent
)

func init() {
	Agent = newAgent()
}

type agentT struct {
	running bool
	timeout time.Duration
	links   atomic.Pointer[[]map[string]string]

	retry retry.AgentT

	ticker   *messaging.Ticker
	emissary *messaging.Channel

	retriever *retrieval.Interface
	requester *request.Interface
	processor template.Agent
}

func newAgent() *agentT {
	a := new(agentT)
	a.retry = retry.NewAgent(timeout)
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
		a.configure(m)
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
func (a *agentT) Link(next core.Exchange) core.Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		ctx, cancel := core.NewContext(nil, a.timeout)
		defer cancel()
		var buf *bytes.Buffer

		switch req.Method {
		case http.MethodGet:
			if req.URL.Path == retrievalPath {
				buf, err = getRetrieval(ctx, a.retriever, a.processor, req)
			} else {
				return httpx.NewResponse(http.StatusBadRequest, nil, nil), errors.New(fmt.Sprintf("resource is invalid [%v] for GET method", req.URL.Path))
			}
		case http.MethodPost:
			switch req.URL.Path {
			case retrievalPath:
				buf, err = postRetrieval(ctx, a.retriever, a.processor, req)
			case relationPath:
				buf, err = relationRetrieval(ctx, a.retriever, a.processor, req)
			case requestThingPath:
				_, err = thingRequest(ctx, a.requester, req.URL.Query())
			case requestLinkPath:
				_, err = linkRequest(ctx, a.requester, req.URL.Query())
			default:
				return httpx.NewResponse(http.StatusBadRequest, nil, nil), errors.New(fmt.Sprintf("resource is invalid [%v]", req.URL.Path))
			}
		default:
			return httpx.NewResponse(http.StatusMethodNotAllowed, nil, nil), nil
		}
		if err != nil {
			return httpx.NewResponse(http.StatusInternalServerError, nil, nil), err
		}
		h := new(http.Header)
		h.Add(httpx.ContentType, httpx.ContentTypeJson)
		return httpx.NewResponse(http.StatusOK, nil, buf), nil
	}
}
