package greetings

import (
	"testing"
)

type Ctx struct {
	name string
	test *testing.T
}

func expectEq(got, want string, ctx Ctx) {
	if got != want {
		ctx.test.Errorf("%s = %s; want %s", ctx.name, got, want)
	}
}

func TestSayHello(t *testing.T) {
	want := "Hello, World!"
	got := SayHello()
	expectEq(got, want, Ctx{"SayHello()", t})
}

func TestSayHelloTo(t *testing.T) {
	want := "Hello, Go!"
	got := SayHelloTo("Go")
	expectEq(got, want, Ctx{"SayHello(Go)", t})
}
