package main

import "fmt"

const a int64 = 10

var b int = 20

type c struct {
	a int
	b int
}

func (c *c) String() string {
	return fmt.Sprintf("%d", c.a*10+c.b)
}

func main() {
	fmt.Println(a, b, c{})
}
