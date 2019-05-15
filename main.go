package main

import (
	"container/list"
	"fmt"
	"github.com/hastechnologyltd/ordis-core/stringutil"
	"github.com/sony/sonyflake"
	"time"
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

	var setupTime = time.Date(2015, 10, 4, 20, 34, 58, 651387237, time.UTC)

	var st sonyflake.Settings
	st.StartTime = setupTime
	fmt.Println(st.StartTime)
	var sf = sonyflake.NewSonyflake(st)
	fmt.Println(sf)

	var id, _ = sf.NextID()
	fmt.Println(id)

	var decomp = sonyflake.Decompose(id)
	fmt.Println(decomp)

	gopherList := list.New()
	gopherList.PushBack(gopher1)
	gopherList.PushBack(gopher2)

	fmt.Println(gopherList.Front().Value)
}

func validateAge(g *gopher) {
	g.isAdult = g.age < 50
}
