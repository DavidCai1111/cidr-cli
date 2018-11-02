package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/DavidCai1993/cidr-cli/cidr"
)

func main() {
	result, err := cidr.Parse(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	b, err := json.Marshal(result)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}
