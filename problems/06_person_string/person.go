package person

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s is %d years old.", p.FirstName, p.LastName, p.Age)
}
