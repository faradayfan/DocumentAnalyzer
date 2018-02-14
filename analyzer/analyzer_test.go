package analyzer

import "testing"

func TestNew(t *testing.T){
	a := New()
	if a == nil {
		t.Error("analyzer is nil")
	}
}