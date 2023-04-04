// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package handler

import (
	"context"
	"net/http"

	"github.com/phelmkamp/gocomps/component"
)

// Props are properties for an HTTP request and response.
type Props struct {
	W http.ResponseWriter
	R *http.Request
}

// New creates a new http.Handler that invokes run on every request.
func New(run func(ctx context.Context, props Props) (child component.Component)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component.Run(r.Context(), component.New(run, Props{W: w, R: r}))
	})
}
