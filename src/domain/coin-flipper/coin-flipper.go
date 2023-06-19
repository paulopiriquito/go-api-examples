package coin_flipper

import (
	"math/rand"
)

func FlipCoin() string {
	var randomResult int
	randomResult = rand.Intn(2)

	if randomResult == 0 {
		return "heads"
	}

	return "tails"
}
