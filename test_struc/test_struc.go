package test_struc

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
	"reflect"

	"github.com/lunixbochs/struc"
)

type Example struct {
	A int `struc:"big"`

	// B will be encoded/decoded as a 16-bit int (a "short")
	// but is stored as a native int in the struct
	B int `struc:"int16"`

	// the sizeof key links a buffer's size to any int field
	Size int `struc:"int8,little,sizeof=Str"`
	Str  string

	// you can get freaky if you want
	Str2 string `struc:"[5]int64,little"`
}

func TestStruc() {
	p := []*int(nil)
	log.Println("p == nil:", p == nil)
	p = []*int{}
	log.Println("p == nil:", p == nil)

	t := Example{1, 2, 0, "test1", "test2"}
	dt, _ := json.MarshalIndent(t, "", "\t")
	log.Println(string(dt))
	log.Println("reflect.ValueOf(t).Kind():", reflect.ValueOf(t).Kind())
	log.Println("reflect.ValueOf(&t).Elem().Kind():", reflect.ValueOf(&t).Elem().Kind())
	log.Println("reflect.ValueOf(&t).Elem().Type():", reflect.ValueOf(&t).Elem().Type())
	log.Println("reflect.ValueOf(t).Field(0).CanSet():", reflect.ValueOf(t).Field(0).CanSet())
	log.Println("reflect.ValueOf(&t).Elem().Field(0).CanSet():", reflect.ValueOf(&t).Elem().Field(0).CanSet())

	var buf bytes.Buffer
	err := struc.Pack(&buf, &t)
	if err != nil {
		return
	}
	log.Println("\n" + hex.Dump(buf.Bytes()))
	o := Example{}
	err = struc.Unpack(&buf, &o)
	if err != nil {
		return
	}
	log.Println(o)
	do, _ := json.MarshalIndent(&o, "", "\t")
	log.Println(string(do))

}
