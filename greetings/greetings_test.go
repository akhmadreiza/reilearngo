package greetings

import (
	"regexp"
	"strings"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Rei"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloNameContains(t *testing.T) {
	name := "Rei"
	msg, err := Hello(name)

	if !strings.Contains(msg, "Rei") || err != nil {
		t.Fatalf(`Hello("Gladys") -> got = %q,  want = contains "Rei", err = %v,`, msg, err)
	}
}

func TestHelloNameEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello ("") -> got = %q, want = should return err, err = %v`, msg, err)
	}

	if err != nil && err.Error() != "name is empty" {
		t.Fatalf(`Hello ("") -> got = "%v", want = "name is empty"`, err)
	}
}