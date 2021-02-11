package mc

import "fmt"

type Operation struct {
	name string
}

func (i *Operation) Execute(content string) {
	fmt.Println("Operation " + i.name + " execute " + content)
}

func NewOperation(name string) *Operation {
	return &Operation{name: name}
}
