// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/phelmkamp/gocomps/component"
)

type props struct {
	in        []int
	out       chan int
	batchSize int
}

func sum(ctx context.Context, p props) component.Component {
	if len(p.in) <= p.batchSize {
		return apply(ctx, p)
	}
	return combine(ctx, p)
}

func split(ctx context.Context, p props) component.Component {
	return component.NewGroup(
		component.New(sum, props{in: p.in[:len(p.in)/2], out: p.out, batchSize: p.batchSize}),
		component.New(sum, props{in: p.in[len(p.in)/2:], out: p.out, batchSize: p.batchSize}),
	)
}

func apply(ctx context.Context, p props) component.Component {
	var n int
	for _, v := range p.in {
		n += v
	}
	p.out <- n
	return component.Component{}
}

func combine(ctx context.Context, p props) component.Component {
	res := make(chan int)
	go func() {
		var n int
		for i := 0; i < 2; i++ {
			n += <-res
		}
		p.out <- n
	}()
	return component.New(split, props{in: p.in, out: res, batchSize: p.batchSize})
}

func main() {
	ints := make([]int, 10_000_000)
	var n int
	for i := 0; i < len(ints); i++ {
		ints[i] = rand.Intn(5)
		n += ints[i]
	}
	fmt.Println("want", n)

	fmt.Println("\nparallel")
	p := props{in: ints, out: make(chan int), batchSize: 100_000}
	start := time.Now()
	component.Run(context.Background(), component.New(sum, p))
	fmt.Println("got", <-p.out, "in", time.Since(start))

	fmt.Println("\nsequential")
	p = props{in: ints, out: make(chan int)}
	start = time.Now()
	go component.Run(context.Background(), component.New(apply, p))
	fmt.Println("got", <-p.out, "in", time.Since(start))
}
