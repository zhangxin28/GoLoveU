package main

import "fmt"

// Liters is a new type which underlining type is float64
type Liters float64

// Milliliters is a new type which underlining type is float64
type Milliliters float64

// Gallons is a new type which underlining type is float64
type Gallons float64

// ToMilliliters converts Liters value to Millitiers
func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

// ToLiters converts Milliliters value to Liters
func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

// RunReceiver tests receiver
func RunReceiver() {
	l := Liters(3)
	fmt.Printf("%.1f liters is %.1f milliliters\n", l, l.ToMilliliters())

	ml := Milliliters(500)
	fmt.Printf("%.1f millilitiers is %.1f liters\n", ml, ml.ToLiters())
}

/*
`receiver` is the value you'are calling the method on.
The `receiver` is listed first when you're calling a method,
and the `receiver` parameter is listed first when you are
defining a method: value this `receiver`, m is the `receiver` paramter
eg:
	type MyType int
	func (m MyType) saiHi() {...}
	value := MyType("a MyType Value")
	value.sayHi()
Once a method is defined on a type, it can be called on any
value of that type.

The `receiver` parameter is (pretty much) just another parameter.
eg:
	func (m MyType) saiHi() {
		fmt.Println("Hi from", m)
	}

By convetion, Go Developers usually use a name consisting of a single
letter: the first letter of the `receiver`'s type name.
Above sample the `receiver` parameter name is `m`.

A method is (pretty much) just like a function
can with return, with parameters, with export/unexport...
eg:
	func (m MyType) WithReturn() int {...}
	func (m MyType) WithParameters(number int, flag bool) int {...}

Pointer receiver paramter can be used to update the receiver itself.
eg:
	func (m *MyType) Double() {
		*m *= 2
	}

Methods with pointer receiver parameter can be called by
both direct values and pointers, because Go autoconverts if needed.
Method with value receiver parameter can be called by
both direct values and pointers, because Go autoconverts if needed.

By convetion, all of your type's methods can take value receivers,
or they can all take pointer receivers.
Do not mixing the two.

To call a method that requires a pointer receiver, you have to
be able to get a pointer to the value!
Below can not work:
	&MyType(2)
	&MyType(2).Double()
You should use a variable to hold it.
eg:
	value := MyType(2)
	value.Double()
	(&value).Double()

Not allowed to define a method on a globally defined type like int/string.
*/
