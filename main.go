package main

import "testing"

func TestSimple(t *testing.T) {
	if true {
		t.Fatal("Expected false,got true")
	}
}

func main() {

}
