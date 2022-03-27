package main

import (
	"learn/disable_redriect"
	"learn/test_c"
	"learn/test_flags"
	"learn/test_json"
	"learn/test_log"
	"learn/test_struc"
	"log"
)

func init() {
	log.SetFlags(log.Flags() | log.Llongfile | log.Lmicroseconds)
}

func main() {
	test_log.TestLog()
	test_flags.TestFlags()
	test_struc.TestStruc()
	disable_redriect.DisableRedirect()
	test_json.TestJsonMarshal()
	test_c.TestC()
}
