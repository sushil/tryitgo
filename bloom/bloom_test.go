package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBloom(t *testing.T) {
	filter := NewBloomFilter(9)

	filter.add("a")
	//filter.add("b")
	//filter.add("c")
	filter.add("d")
	//filter.add("e")
	filter.add("Björn")
	//filter.add("ShoeSeal")

	assert.False(t, filter.exists("Robert"))
	assert.True(t, filter.exists("Björn"))

	filter.setM(1000)
	assert.EqualValues(t, filter.m, 1000)
}