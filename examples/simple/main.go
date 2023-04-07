// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/phelmkamp/gocomps/component"
)

func root(ctx context.Context, props map[string]any) component.Component {
	fmt.Println("root", props)
	return component.NewGroup(
		component.New(a, map[string]any{"name": "a"}),
		component.New(b, map[string]any{"name": "b"}),
	)
}

func a(ctx context.Context, props map[string]any) component.Component {
	fmt.Println("a", props)
	return component.NewGroup(
		component.New(aa, map[string]any{"name": "aa"}),
		component.New(ab, map[string]any{"name": "ab"}),
	)
}

func aa(ctx context.Context, props map[string]any) component.Component {
	fmt.Println("aa", props)
	return component.Component{}
}

func ab(ctx context.Context, props map[string]any) component.Component {
	fmt.Println("ab", props)
	return component.Component{}
}

func b(ctx context.Context, props map[string]any) component.Component {
	fmt.Println("b", props)
	return component.New(ba, map[string]any{"name": "ba"})
}

func ba(ctx context.Context, props map[string]any) component.Component {
	fmt.Println("ba", props)
	return component.Component{}
}

func main() {
	root := component.New(root, map[string]any{"name": "root"})
	component.Run(context.Background(), root)
}
