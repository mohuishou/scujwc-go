package scujwc

import (
	"testing"
)

func TestGPA(t *testing.T) {
	g, err := jwcTest.GPA()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(g)
}

func TestGPAAll(t *testing.T) {
	g, err := jwcTest.GPAAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(g)
}

func TestGPANotPass(t *testing.T) {
	g, err := jwcTest.GPANotPass()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(g)
}
