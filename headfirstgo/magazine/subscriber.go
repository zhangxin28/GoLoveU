package magazine

import "fmt"

/*
A defined type's name must be capitalized if you want to
export it from its package.
Struct field names must be capitalized if you want to
export them from their package.
Even if a struct type is exported from a package,
its fields will be unexported if their names don't
begin with a capital letter.
*/

// Subscriber is a struct
type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

/*
Functions receive a copy of the arguments they're called with,
even for structs.
If you pass a big struct with a lot of fields,
this takes up a lot of the computer's memory:
the original struct and the copy.
It is often a good idea to pass functions a pointer to a struct,
rather than the struct itself.
When you pass a struct pointer, only one copy of the original struct
exists in memory.
The function just receives the momory address
of that single struct, and can read the struct, modify it,
or whatever else it needs to to, all without making an extra copy.
*/

// RunSubScriber shows how to use a struct
func RunSubScriber() {
	s1 := defaultSubscriber("Aman Singh")
	s1.Rate = 4.99
	applyDiscount(s1)
	printInfo(s1)

	s2 := defaultSubscriber("Beth Ryan")
	printInfo(s2)

	// use short variable declaration and omit some some fields.
	// Omitted fields get set to their zero value.
	s3 := &Subscriber{Name: "Bob", Rate: 4.99}
	printInfo(s3)
}

func printInfo(s *Subscriber) {
	/*
		The dot natation to access fields works on struct pointers
		as well as the struct themselves
		Having the write `(*pointer).myField` all the time would get
		tedious quickly,though.
		For this reason, the dot operation lets you access fields via
		pointers to structs, just as you can access fields directly
		from struct values.
		You can leave off the parentheses and the * operation.
	*/
	fmt.Println("Name: ", (*s).Name)
	fmt.Println("Name: ", s.Name)
	fmt.Println("Monthly rate: ", s.Rate)
	fmt.Println("Active: ", s.Active)
}

func defaultSubscriber(name string) *Subscriber {
	return &Subscriber{
		Name:   name,
		Rate:   5.99,
		Active: true,
	}
}

func applyDiscount(s *Subscriber) {
	// Assign to a struct field through the pointer.
	// (*s).rate = 4.99
	s.Rate = 4.99
}
