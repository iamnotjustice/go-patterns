package creationmethod

type Person struct {
	Name string
	Age  int
}

func NewPerson(age int, name string) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
