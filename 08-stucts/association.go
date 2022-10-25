package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Org  *Organziation
}

type Organziation struct {
	Name string
	City string
}

func main() {
	/*
		vmware := Organziation{Name: "VMware", City: "Bangalore"}
		magesh := Employee{Id: 100, Name: "Magesh", Org: &vmware}

		john := Employee{Id: 101, Name: "John", Org: &vmware}
	*/

	vmware := &Organziation{Name: "VMware", City: "Bangalore"}
	magesh := Employee{Id: 100, Name: "Magesh", Org: vmware}

	john := Employee{Id: 101, Name: "John", Org: vmware}

	fmt.Println(magesh)
	fmt.Println(john)

	vmware.City = "Newyork"
	fmt.Printf("%p\n", &vmware)
	fmt.Printf("%p\n", magesh.Org)
	fmt.Printf("%p\n", john.Org)

	fmt.Println(*magesh.Org)
	fmt.Println(*john.Org)

	//dereferencing a pointer to access the fields (DO NOT HAVE TO DO)
	fmt.Println((*vmware).Name)

	//accessing the filed using a pointer to the struct
	fmt.Println(vmware.Name) //equivalent to "vmware -> Name" in c/c++

}
