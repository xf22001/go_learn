package test_c

import "log"

// #cgo CFLAGS: -g -Wall -I.
// #cgo LDFLAGS:
// #include "test_c.h"
import "C"

func TestC() {
	log.Println("add(1, 2)", C.add(1, 2))
}
