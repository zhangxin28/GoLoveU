package magazine

/*
Adding a struct field that is itself a struct type
is no different than adding a field of any other type.
You provide a name for the field, followed by the field's type.
*/

// Employee is a struct
type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

// Address is a struct
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

/*
Go allows you to define anonymous fields: struct fields
that have no name of their own, just a type.
You can use an anonymous field to make our inner struct
easier to access.
When you declare an anonymous field, you can use the field's
type name as if it were the name of the field.
An inner struct that is stored within an outer struct using
an anonymous field is said to be embedded within the outer struct.
Fields for an embedded struct are promoted to the outer struct,
meaning you can access them as if they belong to the outer struct.
But the embedded fields cannot be used in struct literal.
eg:
stuff := Stuff { Name: "" }
fmt.Println(stuff.City)

Bellow code cannot be compiled
stuff1 := Stuff { Name: "", City: ""}
*/

// Stuff is a struct
type Stuff struct {
	Name string
	Address
}
