package main

import (
	"fmt"
	"goloveu/headfirstgo/magazine"
)

// RunMagazine does something with struct
func RunMagazine() {
	magazine.RunSubScriber()
	address := magazine.Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "68111",
	}
	fmt.Println(address)
	subscriber1 := magazine.Subscriber{
		Name:        "Aman Singh",
		HomeAddress: address,
	}
	fmt.Printf("%v\n", subscriber1)
	subscriber2 := magazine.Subscriber{
		Name: "Aman Ryan",
		//HomeAddress: address,
	}
	fmt.Printf("%#v\n", subscriber2)
	fmt.Println(subscriber1.Name, "lives in",
		subscriber1.HomeAddress.City)

	stuff := magazine.Stuff{
		Name:    "Joe",
		Address: address,
	}
	fmt.Println(stuff.Name, "lives in",
		stuff.Address.City)
	fmt.Println(stuff.Name, "lives in",
		stuff.City)
	stuff1 := magazine.Stuff{
		Name: "Joe",
	}
	stuff1.City = "NY"
	fmt.Println(stuff1.Name, "lives in",
		stuff1.City)

	var stuff3 magazine.Stuff
	fmt.Printf("stuff3 = %#v, address = %p\n", stuff3, &stuff3)
	var intvalue int
	fmt.Printf("intvalue = %d, address = %p\n", intvalue, &intvalue)
}
