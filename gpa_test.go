package scujwc

import (
	"testing"
)

func TestGPA(t *testing.T) {
	j, err := NewJwc(2014141453066, "lailin123")
	if err != nil {
		t.Fatal(err)
	}
	grade, err := j.GPA()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(grade)
}
