// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains tests for the buildtag checker.

// ERRORNEXT "possible malformed [+]build comment"
// +builder
//go:build !ignore && toolate
// +build !ignore,toolate

package testdata

// ERRORNEXT "misplaced \+build comment"

var _ = 3

var _ = `
// +build notacomment
`
