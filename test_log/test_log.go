package test_log

import (
	"log"
	"os"
)

func TestLog() {
	logger1 := log.New(os.Stdout, "xiaofei", 0)
	logger2 := log.New(os.Stdout, "xiaofei", log.Llongfile)
	logger3 := log.New(os.Stdout, "xiaofei", log.Llongfile|log.Lmicroseconds)
	logger1.Printf("logger1")
	logger2.Printf("logger2")
	logger3.Printf("logger3")
}
