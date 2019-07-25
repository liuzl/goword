package main

import (
	"testing"
)

func TestNgram(t *testing.T) {
	lines := []string{
		`雄安位于河北省`,
		`天津大学是一所好大学`,
		`what a good idea!`,
	}
	for _, line := range lines {
		ret := Ngram(line)
		for _, r := range ret {
			t.Log(r)
		}
		t.Log("\n")
	}
}
