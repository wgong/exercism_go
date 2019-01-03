package clock

import "fmt"

func ExampleClock_new() {
	// a new clock
	clock1 := New(10, 30)
	fmt.Println(clock1.String())

	// Output: 10:30
}

func ExampleClock_Add() {
	// create a clock
	clock := New(10, 30)

	// add 30 minutes to it
	c := clock.Add(30)
	fmt.Println(c.String())

	// Output: 11:00
}

func ExampleClock_Subtract() {
	// create a clock
	clock := New(10, 30)

	// subtract an hour and a half from it
	c := clock.Subtract(90)
	fmt.Println(c.String())

	// Output: 09:00
}

func ExampleClock_compare() {
	// a new clock
	clock1 := New(10, 30)

	// a second clock, same as the first
	clock2 := New(10, 30)

	// are the clocks equal?
	fmt.Println(clock1 == clock2)

	// change the second clock
	c2 := clock2.Add(30)

	// are the clocks equal now?
	fmt.Println(clock1 == c2)

	// Output:
	// true
	// false
}
