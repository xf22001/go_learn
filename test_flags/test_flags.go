package test_flags

import (
	"flag"
	"log"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// go run . -list1 xiao -list1 fei -list1 "xiao fei"
var myFlags arrayFlags

func TestFlags() {
	flag.Var(&myFlags, "list1", "Some description for this param.")
	flag.Parse()

	log.Println("flag myFlags:", myFlags)
	for i, flagItem := range myFlags {
		log.Println("i:", i, "flagItem:", flagItem)
	}
}
