package reader

import "testing"

func TestNew(t *testing.T) {
	r := New()
	if r == nil{
		t.Error("Reader is nil")
	}
}