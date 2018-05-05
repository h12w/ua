package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"

	"h12.io/ua"
)

func main() {
	dev, err := ua.Parse(os.Args[1], true)
	if err != nil {
		log.Fatal(err)
	}
	buf, err := json.Marshal(dev)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}
