package main

import "testing"

func TestGenerateHash(t *testing.T) {
	salt = "abc"
	expected := "a107ff634856bb300138cac6568c0f2"
	actual := generateHash(0)
	if actual != expected {
		t.Errorf("Want %s, got %s", expected, actual)
	}
}
