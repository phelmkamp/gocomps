// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
	"sync/atomic"
)

const (
	keyComp     ctxKey = "comp"
	keyStateCtr ctxKey = "stateCtr"
)

type ctxKey string

func withStateCtr(ctx context.Context) context.Context {
	ctr := &atomic.Int32{}
	ctr.Store(-1)
	return context.WithValue(ctx, keyStateCtr, ctr)
}

func stateCtr(ctx context.Context) int {
	ctr, _ := ctx.Value(keyStateCtr).(*atomic.Int32)
	return int(ctr.Add(1))
}
