// Copyright 2023 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package crud

// ReadProps are properties for reading a value.
type ReadProps[T any] struct {
	Return func(T)
}

// WriteProps are properties for writing and reading a value.
type WriteProps[T any] struct {
	ReadProps[T]
	V T
}

// NewReadProps creates a new ReadProps.
func NewReadProps[T any](onRet func(T)) ReadProps[T] {
	return ReadProps[T]{Return: onRet}
}

// NewWriteProps creates a new WriteProps.
func NewWriteProps[T any](v T, onRet func(T)) WriteProps[T] {
	return WriteProps[T]{ReadProps: NewReadProps(onRet), V: v}
}
