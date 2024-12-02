package solana

import "testing"

func TestSystemProgramTransfer(t *testing.T) {
	ix := SystemProgramInstructions().Transfer(
		MustParsePubkey("5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrQ"), MustParsePubkey("BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY"), 500_000_000,
	)

	if ix.Data[0] != 2 {
		t.Fatal("Unexpected data")
	}
}
