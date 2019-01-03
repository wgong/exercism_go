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
	pClock := New(10, 30)

	// add 30 minutes to it
	c := (*pClock).Add(30)
	fmt.Println(c.String())

	// Output: 11:00
}

func ExampleClock_Subtract() {
	// create a clock
	pClock := New(10, 30)

	// subtract an hour and a half from it
	c := (*pClock).Subtract(90)
	fmt.Println(c.String())

	// Output: 09:00
}

func ExampleClock_compare() {
	// a new clock
	pClock1 := New(10, 30)

	// a second clock, same as the first
	pClock2 := New(10, 30)

	// are the clocks equal?
	fmt.Println(*pClock1 == *pClock2)

	// change the second clock
	pc2 := (*pClock2).Add(30)

	// are the clocks equal now?
	fmt.Println(*pClock1 == pc2)

	// Output:
	// true
	// false
}
