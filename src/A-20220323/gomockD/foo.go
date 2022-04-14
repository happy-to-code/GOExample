package gomockD

type Foo interface {
	Bar(x int) int
}

func Baz(foo Foo, x int) int {
	return foo.Bar(x) + 1
}
