package scujwc

import (
	"testing"
)

func Test_1(t *testing.T) {
	// var j Jwc
	err := j.Init(2014141453066, "lailin123")
	if err != nil {
		t.Fatal(err)
	}
	j.ScheduleIcal()
}
