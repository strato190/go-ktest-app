package server

import (
	"math/rand"
	"time"
)

func preStopSleepHook() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) // n will be between 0 and 10
	time.Sleep(time.Duration(n) * time.Second)
}
