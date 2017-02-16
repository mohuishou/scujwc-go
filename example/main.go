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
	j.Project()
}
