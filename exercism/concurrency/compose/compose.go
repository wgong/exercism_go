package main

import "fmt"

func doubler(x int) int {
    return 2*x
}

func squarer(x int) int {
    return x*x
}

func Compose(f, g func(x int) int) func(x int) int {
	return func(x int) int {
		return f(g(x))
	}
}
func main() {
	x := 3
	y := Compose(squarer,doubler)(x)
	fmt.Printf("(2*%d)^2 = %d\n", x, y)

	y = Compose(doubler,squarer)(x)
	fmt.Printf("(%d)^2 * 2 = %d\n", x, y)



}