package lib

import "testing"

func TestAdd(t *testing.T) {
	v := Add(1, 2)
	if v != 3 {
		t.Errorf("Got %d, expected 3.", v)
	}
}
