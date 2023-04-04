// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package crud

// ReadProps are properties for reading a value.
type ReadProps[T any] struct {
	OnReturn func(T)
	OnError  func(error)
}

// WriteProps are properties for writing and reading a value.
type WriteProps[T any] struct {
	ReadProps[T]
	V T
}

// NewReadProps creates a new ReadProps.
func NewReadProps[T any](onRet func(T), onErr func(error)) ReadProps[T] {
	return ReadProps[T]{OnReturn: onRet, OnError: onErr}
}

// NewWriteProps creates a new WriteProps.
func NewWriteProps[T any](v T, onRet func(T), onErr func(error)) WriteProps[T] {
	return WriteProps[T]{ReadProps: NewReadProps(onRet, onErr), V: v}
}
