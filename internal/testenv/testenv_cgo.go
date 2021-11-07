// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo

package testenv

// https://github.com/golang/go/blob/9e6ad46bccfa7a63e768236bcd1fd54dab38e4d1/src/internal/testenv/testenv_cgo.go

func init() {
	haveCGO = true
}
