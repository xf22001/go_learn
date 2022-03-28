package test_syntax

import (
	"log"
	"time"
)

func test_select() {
	duration := time.Second * 1
	timer := time.NewTimer(duration)
	defer timer.Stop()

	for i := 0; i < 10; i++ {

		select {
		case <-timer.C:
			log.Println("timeout")
			timer.Reset(time.Second * 1)
		}
	}
}

func TestSyntax() {
	test_select()
}
