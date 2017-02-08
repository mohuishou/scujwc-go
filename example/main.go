package main

import (
	"fmt"
	"os"

	"encoding/json"

	"../"
)

func main() {
	s := scujwc.Jwc{}
	s.Init(2014141453066, "lailin123")
	_ = s.Login()
	grade, e := s.GPA()
	if e != nil {
		return
	}
	j, err := json.Marshal(grade)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(j)

}
