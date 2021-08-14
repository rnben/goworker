package goworker

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func TestWorker(t *testing.T) {
	q := NewQueen(20)

	for i := 0; i < 100; i++ {
		q.Add(func() {
			time.Sleep(time.Second)
		})
	}

	q.Run()
}
