package main

import "testing"

func TestResponseMsg(t *testing.T) {
	if responseMsg() != "pong" {
		t.Fatal("Reponse msg has to be pong.")
	}
}
