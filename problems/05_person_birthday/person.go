package person

type Person struct {
	Name string
	Age  int
}

func (p *Person) Birthday() {
	p.Age++
}
