package component

import (
	"context"
)

var states = make(map[uintptr][]any)

// UseState adds a state variable for the current Component.
func UseState[T any](ctx context.Context, initial T) (current T, setter func(T)) {
	comp, _ := ctx.Value(keyComp).(Component)
	ctr := stateCtr(ctx)
	if ctr == len(states[comp.id]) {
		states[comp.id] = append(states[comp.id], initial)
	}
	return states[comp.id][ctr].(T), func(v T) {
		states[comp.id][ctr] = v
		go Run(ctx, comp)
	}
}
