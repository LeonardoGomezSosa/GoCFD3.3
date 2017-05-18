package main

import (
	"./XMLStructures"
)

func main() {
	var comprobante xmlstructures.Comprobante
	xmlstructures.MarshallData2XML(comprobante)

}
