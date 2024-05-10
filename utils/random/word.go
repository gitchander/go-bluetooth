package random

import (
	"math/rand"
)

//------------------------------------------------------------------------------
// const bitsPerWord = 32

// type word uint32

// func randWord(r *rand.Rand) word {
// 	return word(r.Uint32())
// }

//------------------------------------------------------------------------------

const bitsPerWord = 64

type word uint64

func randWord(r *rand.Rand) word {
	return word(r.Uint64())
}

//------------------------------------------------------------------------------
