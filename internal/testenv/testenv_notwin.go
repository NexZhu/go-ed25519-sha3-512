// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

package testenv

// https://github.com/golang/go/blob/9e6ad46bccfa7a63e768236bcd1fd54dab38e4d1/src/internal/testenv/testenv_notwin.go

import (
	"runtime"
)

func hasSymlink() (ok bool, reason string) {
	switch runtime.GOOS {
	case "android", "plan9":
		return false, ""
	}

	return true, ""
}
