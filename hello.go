package main

import (
	"errors"
	"fmt"
	"github.com/AJSMac/newmath"
)

func init() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)

	s2 = "C" + s[2:4]
	fmt.Printf("%s\n", s2[0:3])

	fmt.Println("Hello World. Sqrt(2)=", newmath.Sqrt(2))

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)

		const (
			x = iota
			y = iota
			z = iota
		)
		const v = iota

		test1 := make(map[string]int)
		test2 := new(int)
		fmt.Println(test1, " ", test2)
		numbers := make(map[string]int)
		fmt.Println(*test2)
		*test2 = 5
		fmt.Println(*test2)
		numbers["one"] = 1
		numbers["ten"] = 10
		numbers["three"] = 3
		fmt.Println(test2)

		fmt.Println("the third number is : ", numbers["three"])

	}

}
