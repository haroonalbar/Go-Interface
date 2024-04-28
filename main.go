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

  // if it's an empty interface it can hold any value 
  var j interface{} = "boi"

  // we can get the value of underlying type j if value not preset give it's zero value and ok would be flase
  k ,ok := j.(string)
  fmt.Println(k,ok)

  bro,ok := j.(int)
  fmt.Println(bro,ok)


	var t *Car
	i = t

	a, err := i.car()
	fmt.Println(a, "\t", err)

	i = &Car{"Bugatti"}
	fmt.Println(i.car())

}
