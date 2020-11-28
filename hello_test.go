package darktheatre

import "testing"

func TestDtheatre(t *testing.T) {
    want := "Hello, D."
    if got := Dtheatre(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}