package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()

	//image input
	client.SetImage("ocr-test.jpg")

	//define language
	client.SetLanguage("eng")

	//output
	output, err := client.Text()
	if err != nil {
		panic("failed to extract the text")
	}

	fmt.Println(output)
}
