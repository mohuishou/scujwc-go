package scujwc

import (
	"testing"
)

func TestGet(t *testing.T) {
	userinfo, err := jwcTest.UserInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userinfo)
}
