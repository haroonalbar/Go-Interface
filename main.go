package main

import (
	"errors"
	"fmt"
)

type Inter interface {
	car() (string, error)
}

type Car struct {
	name string
}

func (c *Car) car() (string, error) {
	if c == nil {
		return "nil", errors.New("nil pointer")
	}
	return c.name, nil
}

func main() {
	var i Inter
	var t *Car
	i = t

	a, err := i.car()
	fmt.Println(a, "\t", err)

	i = &Car{"Bugatti"}
	fmt.Println(i.car())

}
