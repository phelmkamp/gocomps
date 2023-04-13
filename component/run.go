// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
)

var queue = make(chan job)

type job struct {
	ctx  context.Context
	comp Component
	done chan struct{}
}

// Run walks the tree made up of the given Component plus its descendants and runs each one.
func Run(ctx context.Context, comp Component) {
	// add job to queue and wait for it to complete
	done := make(chan struct{})
	queue <- job{ctx: ctx, comp: comp, done: done}
	<-done
}

func run(ctx context.Context, comp Component) {
	if comp.run == nil {
		// nothing to do
		return
	}

	if ctx.Err() != nil {
		// context is done
		return
	}

	ctx = context.WithValue(ctx, keyComp, comp)
	ctx = withStateCtr(ctx)
	child := comp.run(ctx, comp.props)
	run(ctx, child)
}

func init() {
	// start background goroutine to run jobs in queue
	go func() {
		for job := range queue {
			run(job.ctx, job.comp)
			job.done <- struct{}{}
		}
	}()
}
