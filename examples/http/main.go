// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/phelmkamp/gocomps/component"
	"github.com/phelmkamp/gocomps/handler"
)

type greetProps struct {
	name  string
	greet func(msg string, err error)
}

func handle(ctx context.Context, props handler.Props) component.Component {
	name := props.R.URL.Query().Get("name")
	onGreet := func(msg string, err error) {
		if err != nil {
			props.W.WriteHeader(http.StatusBadRequest)
			props.W.Write([]byte(err.Error()))
			return
		}
		props.W.Write([]byte(msg))
	}
	return component.New(greetSvc, greetProps{name: name, greet: onGreet})
}

func greetSvc(ctx context.Context, props greetProps) component.Component {
	if props.name == "" {
		props.greet("", errors.New("name not provided"))
		return component.Component{}
	}
	props.greet("hello, "+props.name, nil)
	return component.Component{}
}

func main() {
	http.ListenAndServe(":8080", handler.New(handle))
}
