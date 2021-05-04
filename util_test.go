// Copyright 2021 Patrick Zhong.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringFirstUpper(t *testing.T) {
	assert.Equal(t, "A", stringFirstUppper("a"))
	assert.Equal(t, "B", stringFirstUppper("b"))
	assert.Equal(t, "C", stringFirstUppper("c"))
	assert.Equal(t, "D", stringFirstUppper("d"))
	assert.Equal(t, "E", stringFirstUppper("e"))
	assert.Equal(t, "F", stringFirstUppper("f"))
	assert.Equal(t, "G", stringFirstUppper("g"))
	assert.Equal(t, "H", stringFirstUppper("h"))
	assert.Equal(t, "I", stringFirstUppper("i"))
	assert.Equal(t, "J", stringFirstUppper("j"))
	assert.Equal(t, "K", stringFirstUppper("k"))
	assert.Equal(t, "L", stringFirstUppper("l"))
	assert.Equal(t, "M", stringFirstUppper("m"))
	assert.Equal(t, "N", stringFirstUppper("n"))
	assert.Equal(t, "O", stringFirstUppper("o"))
	assert.Equal(t, "P", stringFirstUppper("p"))
	assert.Equal(t, "Q", stringFirstUppper("q"))
	assert.Equal(t, "R", stringFirstUppper("r"))
	assert.Equal(t, "T", stringFirstUppper("t"))
	assert.Equal(t, "S", stringFirstUppper("s"))
	assert.Equal(t, "X", stringFirstUppper("x"))
	assert.Equal(t, "Y", stringFirstUppper("y"))
	assert.Equal(t, "Z", stringFirstUppper("z"))

	assert.Equal(t, "A", stringFirstUppper("A"))
	assert.Equal(t, "B", stringFirstUppper("B"))
	assert.Equal(t, "C", stringFirstUppper("C"))
	assert.Equal(t, "D", stringFirstUppper("D"))
	assert.Equal(t, "E", stringFirstUppper("E"))
	assert.Equal(t, "F", stringFirstUppper("F"))
	assert.Equal(t, "G", stringFirstUppper("G"))
	assert.Equal(t, "H", stringFirstUppper("H"))
	assert.Equal(t, "I", stringFirstUppper("I"))
	assert.Equal(t, "J", stringFirstUppper("J"))
	assert.Equal(t, "K", stringFirstUppper("K"))
	assert.Equal(t, "L", stringFirstUppper("L"))
	assert.Equal(t, "M", stringFirstUppper("M"))
	assert.Equal(t, "N", stringFirstUppper("N"))
	assert.Equal(t, "O", stringFirstUppper("O"))
	assert.Equal(t, "P", stringFirstUppper("P"))
	assert.Equal(t, "Q", stringFirstUppper("Q"))
	assert.Equal(t, "R", stringFirstUppper("R"))
	assert.Equal(t, "T", stringFirstUppper("T"))
	assert.Equal(t, "S", stringFirstUppper("S"))
	assert.Equal(t, "X", stringFirstUppper("X"))
	assert.Equal(t, "Y", stringFirstUppper("Y"))
	assert.Equal(t, "Z", stringFirstUppper("Z"))
}
