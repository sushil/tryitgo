package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math/big"
)

type BloomFilter struct {
	m        int64
	elements []bool
	filters  []func(elem string) string
}

func NewBloomFilter(size int64) BloomFilter {
	return BloomFilter{
		m:        size,
		elements: make([]bool, size),
		filters: []func(elem string) string{
			FilterMD5,
			FilterSHA1,
		},
	}
}

func (bf BloomFilter) setM(m int64) {
	bf.m = m
}

func (bf BloomFilter) exists(elem string) bool {
	for _, filter := range bf.filters {
		h := filter(elem)
		pos := bf.mapHashToPos(h)

		if !bf.elements[pos] {
			return false
		}
	}

	return true
}

func (bf BloomFilter) add(elem string) {
	for _, filter := range bf.filters {
		h := filter(elem)
		pos := bf.mapHashToPos(h)

		bf.elements[pos] = true
	}
}

func (bf BloomFilter) mapHashToPos(h string) int64 {
	bi := big.NewInt(0)
	bi.SetString(h, 16)
	z := big.NewInt(0)
	z.Mod(bi, big.NewInt(bf.m))
	return z.Int64()
}

func FilterMD5(elem string) string {
	h := md5.New()
	h.Write([]byte(elem))
	return hex.EncodeToString(h.Sum(nil))
}

func FilterSHA1(elem string) string {
	sha := sha1.New()
	sha.Write([]byte(elem))
	return hex.EncodeToString(sha.Sum(nil))
}
