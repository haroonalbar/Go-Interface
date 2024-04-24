package main

import "fmt"

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

func main() {
  var i Inter = &Car{"Bugatti"}
  fmt.Println(i.car())
}
