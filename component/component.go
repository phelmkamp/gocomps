// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
)

// Component is a reusable piece of functionality.
type Component struct {
	run   func(ctx context.Context, props any) (child Component)
	props any
}

// RunFunc is the functionality encapsulated by a Component.
type RunFunc[P any] func(ctx context.Context, props P) (child Component)

// New creates a new Component.
func New[P any](run RunFunc[P], props P) Component {
	return Component{
		run:   func(ctx context.Context, props any) (child Component) { return run(ctx, props.(P)) },
		props: props,
	}
}

// NewGroup creates a new group of the given Components.
// This is useful when a Component has more than one child.
func NewGroup(comps ...Component) Component {
	return Component{
		run: func(ctx context.Context, props any) Component {
			for _, c := range comps {
				Run(ctx, c)
			}
			return Component{}
		},
	}
}

// Run walks the tree made up of the given Component plus its descendants and runs each one.
func Run(ctx context.Context, comp Component) {
	if comp.run != nil {
		child := comp.run(ctx, comp.props)
		Run(ctx, child)
	}
}
