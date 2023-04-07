// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
	"reflect"
	"sync"
)

// Component is a reusable piece of functionality.
type Component struct {
	id    uintptr
	run   func(ctx context.Context, props any) (child Component)
	props any
}

// RunFunc is the functionality encapsulated by a Component.
type RunFunc[P any] func(ctx context.Context, props P) (child Component)

// New creates a new Component.
func New[P any](run RunFunc[P], props P) Component {
	return Component{
		id: reflect.ValueOf(run).Pointer(),
		run: func(ctx context.Context, props any) (child Component) {
			p, _ := props.(P)
			return run(ctx, p)
		},
		props: props,
	}
}

// NewGroup creates a new group of the given Components.
// This is useful when a Component has more than one child.
func NewGroup(comps ...Component) Component {
	return Component{
		run: func(ctx context.Context, props any) Component {
			wg := &sync.WaitGroup{}
			wg.Add(len(comps))
			for i := range comps {
				go func(c Component) {
					defer wg.Done()
					run(ctx, c)
				}(comps[i])
			}
			wg.Wait()
			return Component{}
		},
	}
}
