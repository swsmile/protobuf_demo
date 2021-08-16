package main

import (
	"example.com/protobuf_demo/gen_proto/example.com/protobuf_demo/go/sw_proto"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	book := &sw_proto.AddressBook{}
	// ...

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("myFile", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
