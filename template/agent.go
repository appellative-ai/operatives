package template

import (
	"errors"
	"github.com/appellative-ai/common/messaging"
	//"github.com/appellative-ai/core/std"
	"github.com/appellative-ai/postgres/retrieval"
	"time"
)

const (
	AgentName  = "common:core:agent/template/center"
	timeout    = time.Second * 4
	defaultSql = "CALL dbo.Representation($1)"
)

type Agent interface {
	messaging.Agent
	Add(entry Entry)
	Build(name string, args []Arg) (Result, error)
}

type agentT struct {
	timeout   time.Duration
	cache     *std.MapT[string, Entry]
	retriever *retrieval.Interface
}

func NewAgent(retriever *retrieval.Interface) Agent {
	return newAgent(retriever)
}

func newAgent(retriever *retrieval.Interface) *agentT {
	a := new(agentT)
	a.timeout = timeout
	a.cache = std.NewSyncMap[string, Entry]()
	a.retriever = retriever
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
		messaging.UpdateContent[time.Duration](m, &a.timeout)
		return
	}
}

func (a *agentT) Add(entry Entry) {
	if entry.Name == "" {
		return
	}
	a.cache.Store(entry.Name, entry)
}

func (a *agentT) Build(name string, args []Arg) (Result, error) {
	if name == "" {
		return Result{}, errors.New("name is empty")
	}
	if len(args) == 0 {
		return Result{}, errors.New("arguments are empty")
	}
	t, ok := a.cache.Load(name)
	if !ok {
		var err error
		t, err = a.add(name)
		if err != nil {
			return Result{}, err
		}
	}
	newArgs, err := Build(args, t.Params)
	if err != nil {
		return Result{}, err
	}
	return Result{Sql: t.Sql, Args: newArgs}, nil
}
