package fbutils

import (
	"math/rand"
	"time"
	"fmt"
)

// Picks a random integer using rand.Intn(n) seeded with the time.
func PickRand(n int) int {
	rand.Seed(time.Now().UnixNano())
	pick := rand.Intn(n)
	return pick
}

// Takes a uint32 and returns a string
// of that integer zero-padded to 5 digits.
func PadNumberWithZero(value uint32) string {
	return fmt.Sprintf("%05d", value)
}


