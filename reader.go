package main

import (
	"example.com/protobuf_demo/gen_proto/example.com/protobuf_demo/go/sw_proto"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	// Read the existing address book.
	in, err := ioutil.ReadFile("myFile")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &sw_proto.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
}
