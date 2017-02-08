package main

import (
	"fmt"

	"../"
)

func main() {
	s := scujwc.Jwc{}
	s.Init(2014141453066, "lailin123")
	err := s.Login()
	e := s.GPA()
	fmt.Println(err, e)
}
