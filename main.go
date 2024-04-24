package main

import (
	"fmt"
	"strings"
)

type Inter interface {
  car() string
}

type Car struct{
  name string
}

//here type Car implements interface Inter but we don't have to specify it cause it's defined implicitly
func (c *Car) car() string{
  return c.name
}

type NewString string

// here we have a new method with same name but diffrent underlying type of NewString
func (i NewString) car() string{
  up := strings.ToUpper(string(i))
  return up
}

func main() {
  var i Inter 
  i = &Car{"Bugatti"}
  fmt.Println(i.car())

  //here the the underlying method type changed so interface opts for that method.
  i = NewString(i.car())
  fmt.Println(i.car())
}
