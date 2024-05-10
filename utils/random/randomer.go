package random

import (
	"math/rand"
)

const bitsPerByte = 8

type Bit uint

type Randomer struct {
	rd *rand.Rand

	x word // bits accumulator
	n int  // bits available
}

func NewRandomer(rd *rand.Rand) *Randomer {
	return &Randomer{
		rd: rd,
	}
}

func (r *Randomer) Rand() *rand.Rand {
	return r.rd
}

func (r *Randomer) randAccum() {
	r.x = randWord(r.rd)
	r.n = bitsPerWord
}

func (r *Randomer) Bool() bool {

	const size = 1

	if r.n < size {
		r.randAccum()
	}

	b := ((r.x & 1) == 1)

	r.x >>= size
	r.n -= size

	return b
}

func (r *Randomer) Bit() Bit {

	const size = 1

	if r.n < size {
		r.randAccum()
	}

	t := Bit(r.x & 1)

	r.x >>= size
	r.n -= size

	return t
}

func (r *Randomer) Byte() byte {

	const size = bitsPerByte

	if r.n < size {
		r.randAccum()
	}

	b := byte(r.x)

	r.x >>= size
	r.n -= size

	return b
}

func (r *Randomer) FillBytes(bs []byte) {

	const size = bitsPerByte

	for i := range bs {

		if r.n < size {
			r.randAccum()
		}

		bs[i] = byte(r.x)

		r.x >>= size
		r.n -= size
	}
}
