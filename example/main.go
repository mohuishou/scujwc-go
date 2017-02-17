package main

import (
	"fmt"

	"../"
)

func main() {
	var j scujwc.Jwc
	err := j.Init(2014141453066, "lailin123")
	if err != nil {
		fmt.Println(err)
	}
	p, err := scujwc.Str2proc("[999008030]中华文化（艺术篇）", "0")

	// err = j.Project()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
