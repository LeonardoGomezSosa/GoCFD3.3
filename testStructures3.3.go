package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"

	"./XMLStructures"
)

func main() {
	var comprobante xmlstructures.Comprobante
	estructura := xmlstructures.MarshallData2XML(comprobante)
	fmt.Println(estructura)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
}
