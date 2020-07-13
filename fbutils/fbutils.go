package fbutils

import (
	"math/rand"
	"time"
	"fmt"
)

func PickRand(n int) int {
	rand.Seed(time.Now().UnixNano())
	pick := rand.Intn(n)
	return pick
}

func PadNumberWithZero(value uint32) string {
	return fmt.Sprintf("%05d", value)
}


