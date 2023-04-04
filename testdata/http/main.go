// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"net/http"

	"github.com/phelmkamp/gocomps/component"
	"github.com/phelmkamp/gocomps/crud"
	"github.com/phelmkamp/gocomps/handler"
)

func handle(ctx context.Context, props handler.Props) component.Component {
	name := props.R.URL.Query().Get("name")
	onGreet := func(s string) {
		props.W.Write([]byte(s))
	}
	return component.New(greetSvc, crud.NewWriteProps(name, onGreet, nil))
}

func greetSvc(ctx context.Context, props crud.WriteProps[string]) component.Component {
	props.OnReturn("hello, " + props.V)
	return component.Component{}
}

func main() {
	http.ListenAndServe(":8080", handler.New(handle))
}
