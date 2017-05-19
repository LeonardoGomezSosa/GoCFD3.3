package main

import (
	"fmt"

	"./XMLStructures"
)

func main() {
	var comprobante xmlstructures.Comprobante
	estructura := xmlstructures.MarshallData2XML(comprobante)
	fmt.Println(estructura)

}
