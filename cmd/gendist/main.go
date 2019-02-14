package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JeanJacquesSerpoul/bridge/tools"
)

func main() {
	fo, err := os.Create("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	if err := tools.GenDistWithPointToFile(fo, 1, 13); err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
}
