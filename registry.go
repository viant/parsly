package parsly

import (
	"github.com/pkg/errors"
	"github.com/viant/parsly/lex"
)

//Registry represents matcher
type Registry struct {
	registry []*Token
}

//Add adds a matcher
func (m *Registry) Add(code int, name string, matcher lex.Matcher) error {
	if code < 0 {
		return errors.Errorf("code has to be positive number, but had: %v", code)
	}

	if code >= len(m.registry) {
		extra := make([]*Token, code-len(m.registry))
		m.registry = append(m.registry, extra...)
	}
	m.registry[code] = &Token{Code: code, Name: name, Matcher: matcher}
	return nil
}

//Lookup returns register matched for supplied code or nil
func (m *Registry) Lookup(ID int) (*Token, bool) {
	token := m.registry[ID]
	return token, true
}

//NewRegistry creates a matcher registry
func NewRegistry() *Registry {
	return &Registry{registry: make([]*Token, 20)}
}
