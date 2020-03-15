package nonce

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New() uint64 {
	return rand.Uint64()
}
