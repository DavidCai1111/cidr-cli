package main

import (
	"fmt"

	"github.com/DavidCai1993/cidr-cli/cidr"
)

func main() {
	_, err := cidr.Parse("192.168.23.35/21")

	fmt.Println(err)
}
