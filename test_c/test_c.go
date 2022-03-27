package test_c

import "log"

// #include "test_c.h"
import "C"

func TestC() {
	log.Println("add(1, 2)", C.add(1, 2))
}

//export sum
func sum(a, b C.int) C.int {
	return a + b
}
