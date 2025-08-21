package template

import "errors"

func (a *agentT) add(name string) (Entry, error) {
	if name == "" {
		return Entry{}, errors.New("name is empty")
	}
	t, err := a.get(name)
	if err != nil {
		return Entry{}, err
	}
	a.cache.Store(name, t)
	return t, nil
}
