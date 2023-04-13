// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package component

import "context"

// ReducerFunc returns the updated state based on the given action.
type ReducerFunc[S, A any] func(state S, action A) (updatedState S)

// DispatchFunc dispatches the given action to update the state.
type DispatchFunc[A any] func(action A)

// UseReducer adds a reducer for the current Component.S
func UseReducer[S, A any](ctx context.Context, reduce ReducerFunc[S, A], initial S) (current S, dispatch DispatchFunc[A]) {
	comp, _ := ctx.Value(keyComp).(Component)
	ctr := stateCtr(ctx)
	if ctr == len(states[comp.id]) {
		states[comp.id] = append(states[comp.id], initial)
	}
	return states[comp.id][ctr].(S), func(action A) {
		states[comp.id][ctr] = reduce(states[comp.id][ctr].(S), action)
		go Run(ctx, comp)
	}
}
