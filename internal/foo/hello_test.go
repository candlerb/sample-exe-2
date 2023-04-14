package foo_test

import "strings"
import "testing"
import "github.com/candlerb/sample-exe-2/internal/foo"

func TestGreeting(t *testing.T) {
	g := foo.Greeting()
	if !strings.Contains(g, "Hello") {
		t.Errorf("Greeting is not polite: %#v", g)
	}
}
