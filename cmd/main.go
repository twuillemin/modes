package main

import (
	"encoding/json"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/reader"
	"log"
)

func main() {

	msg1 := []uint8{0x02, 0xE1, 0x95, 0x30, 0x1A, 0xD7, 0xB8}

	message1, err := reader.ReadMessage(msg1)
	if err != nil {
		log.Fatal(err)
	}
	bytes1, err := json.Marshal(message1)
	if err != nil {
		log.Fatalf("Can't serialize due to %v", err)
	}
	fmt.Printf("Message 1: %v\n", string(bytes1))

	msg2 := []uint8{0x5F, 0x46, 0x90, 0xF9, 0x76, 0xF4, 0xE9}

	message2, err := reader.ReadMessage(msg2)
	if err != nil {
		log.Fatal(err)
	}
	bytes2, err := json.Marshal(message2)
	if err != nil {
		log.Fatalf("Can't serialize due to %v", err)
	}
	fmt.Printf("Message 2: %v\n", string(bytes2))

	fmt.Println()
	message1.PrettyPrint()
	fmt.Println()
	message2.PrettyPrint()
	fmt.Println()

}
