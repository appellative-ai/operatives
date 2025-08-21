package template

import (
	"encoding/json"
	"github.com/appellative-ai/common/iox"
)

func AddEntry(agent Agent, fileName string) error {
	buf, err := iox.ReadFile(fileName)
	if err != nil {
		return err
	}
	var e []Entry
	err = json.Unmarshal(buf, &e)
	if err != nil {
		return err
	}
	for _, entry := range e {
		agent.Add(entry)
	}
	return nil
}
