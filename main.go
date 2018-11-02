package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/DavidCai1993/cidr-cli/cidr"
)

func main() {
	result, err := cidr.Parse("192.168.23.35/21")

	if err != nil {
		log.Fatalln(err)
	}

	b, err := json.Marshal(result)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}
