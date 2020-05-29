package main

// Address :)
type Address struct {
	HouseNumber int
	Street      string
	PostCode    string
	City        string
	Country     string
}

// Person :)
type Person struct {
	Name string
	Age  int
	Home Address
}
