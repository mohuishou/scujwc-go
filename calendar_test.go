package scujwc

import (
	"os"
	"testing"
)

func Test_1(t *testing.T) {
	var j Jwc
	err := j.Init(2014141453066, "lailin123")
	if err != nil {
		t.Fatal(err)
	}
	ical, err := j.Calendar(2)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Create("ical.ics")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()
	f.Write(ical.Bytes())
}
