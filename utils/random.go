package utils

import (
	"math/rand"
	"time"
)

var rCh = make(chan int, 128)

func init() {
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			rCh <- r.Int()
		}
	}()
}

func RInt() int {
	return <-rCh
}
