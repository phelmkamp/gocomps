// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
)

type Component struct {
	run   func(ctx context.Context, props any) (child Component)
	props any
}

func New[P any](run func(ctx context.Context, props P) (child Component), props P) Component {
	return Component{
		run:   func(ctx context.Context, props any) (child Component) { return run(ctx, props.(P)) },
		props: props,
	}
}

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

func Run(ctx context.Context, comp Component) {
	if comp.run != nil {
		child := comp.run(ctx, comp.props)
		Run(ctx, child)
	}
}
