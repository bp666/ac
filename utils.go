package ac

import (
	"math/rand"
	"time"
)

// Sleep n second
func Sleep(sec uint16) {
	time.Sleep(time.Duration(sec) * time.Second)
}

// SleepRand [left, right] second (right > left, left >= 0, right > 0)
func SleepRand(left, right uint16) {
	var sec int
	if right != 0 {
		rand.Seed(time.Now().UnixNano())
		sec = rand.Intn(int(right)) + int(left) + 1
	} else {
		sec = 0
	}
	time.Sleep(time.Duration(sec) * time.Second)
}
