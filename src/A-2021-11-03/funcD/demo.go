package main

type Greeting func(name string) string

func main() {
	var greeting Greeting
	greeting = func(name string) string {
		return name + "hahah"
	}

	println(greeting("yes"))
}
