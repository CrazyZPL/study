package main

import (
	"strings"
	"fmt"
)

func main() {
	text := "zpl"
	text2 := "zpl zpl"
	fmt.Println(strings.Split(text, " "), strings.Split(text2, " "))
}
