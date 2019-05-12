package main

import (
	"fmt"
	"github.com/hastechnologyltd/ordis-core/stringutil"
)

type gopher struct {
	name    string
	age     int
	isAdult bool
}

func (g gopher) jump() string {
	if g.age < 50 {
		return g.name + " can jump high"
	} else {
		return g.name + " can jump low"
	}
}

func main() {

	gopher1 := &gopher{name: "Jeff", age: 47}
	gopher2 := &gopher{name: "Susie", age: 50}

	validateAge(gopher1)
	validateAge(gopher2)

	fmt.Println(gopher1)
	fmt.Println(gopher2)

	fmt.Println(stringutil.Reverse("Jeff"))
}

func validateAge(g *gopher) {
	g.isAdult = g.age < 50
}
