package fhtml

import "context"

type Valuable interface {
	Value(key string) any
}

// ValuableMap extends Valuable
var _ Valuable = (*ValuableMap)(nil)

type ValuableMap map[string]any

func (v ValuableMap) Value(key string) any {
	return v[key]
}

// ValuableContext extends Valuable
var _ Valuable = (*ValuableContext)(nil)

type ValuableContext struct {
	context.Context
}

func (v ValuableContext) Value(key string) any {
	return v.Context.Value(key)
}
