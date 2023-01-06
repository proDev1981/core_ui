package models

type Person struct {
	Name string
	Age  int
}

var Persons = []Person{
	{
		Name: "alberto",
		Age:  41,
	},
	{
		Name: "paco",
		Age:  39,
	},
}
