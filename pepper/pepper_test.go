package pepper

import "testing"

func TestNew(t *testing.T) {
	r := New()
	if r == nil{
		t.Error("Pepper is nil")
	}
}