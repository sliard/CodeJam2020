package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {

	assert := assert.New(t)

	a := []int{0, 1, 1, 1}
	ra := []int{1, 1, 1, 0}
	a = Reverse(a)
	assert.Equal(a, ra)

	b := []int{0, 1, 0, 1, 1, 1, 0}
	rb := []int{0, 1, 1, 1, 0, 1, 0}
	b = Reverse(b)
	assert.Equal(b, rb)
}

func TestBitFlip(t *testing.T) {

	assert := assert.New(t)

	aa := []int{0, 1, 1, 1}
	a := []int{0, 1, 1, 1}
	ra := []int{1, 0, 0, 0}
	ta := BitFlip(a)
	assert.Equal(ta, ra)
	assert.Equal(aa, a)

	b := []int{0, 1, 0, 1, 1, 1, 0}
	rb := []int{1, 0, 1, 0, 0, 0, 1}
	tb := BitFlip(b)
	assert.Equal(tb, rb)
}
