// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/phelmkamp/gocomps/component"
)

func reduce(i int, action string) int {
	switch action {
	case "add":
		return i + 1
	case "subtract":
		return i - 1
	default:
		return i
	}
}

func read(ctx context.Context, done chan struct{}) component.Component {
	i, dispatch := component.UseReducer(ctx, reduce, 0)

	fmt.Println(i)

	onChange := func(s string) {
		if s == "+" {
			dispatch("add")
		} else if s == "-" {
			dispatch("subtract")
		} else {
			done <- struct{}{}
		}
	}
	return component.New(input, onChange)
}

func input(ctx context.Context, onChange func(string)) component.Component {
	go func() {
		r := termbox.PollEvent().Ch
		onChange(string(r))
	}()
	return component.Component{}
}

func main() {
	termbox.Init()
	defer termbox.Close()
	fmt.Println("Enter +/- to add/subtract:")
	done := make(chan struct{})
	component.Run(context.Background(), component.New(read, done))
	<-done
}
