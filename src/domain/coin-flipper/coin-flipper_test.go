package coin_flipper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinFlip_MustReturnHeadsOrTails(t *testing.T) {
	for i := 1; i < 5; i++ {
		result := "nil"
		result = FlipCoin()
		assert.Subset(t, []string{"heads", "tails"}, []string{result})
	}
}
