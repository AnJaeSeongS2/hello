package main

import (
	"fmt"
	"os"
)

func main() {

	openFile("RecodverTest.txt")
	defer println("defer")

	println("Done")
	defer println("defer2")
	defer println("defer3")
}

func openFile(fn string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Open error", r)
		}
	}()
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer fmt.Println("f.close?2")
	defer f.Close()
	defer fmt.Println("f.close?3")

}
