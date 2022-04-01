package main

import (
	"learn/test_regexp"
	"log"
)

func init() {
	log.SetFlags(log.Flags() | log.Llongfile | log.Lmicroseconds)
}

func main() {
	// test_log.TestLog()
	// test_flags.TestFlags()
	// test_struc.TestStruc()
	// disable_redriect.DisableRedirect()
	// test_json.TestJsonMarshal()
	// test_c.TestC()
	// test_syntax.TestSyntax()
	test_regexp.TestRegexp()
}
