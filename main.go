package main

import (
	"fmt"
	"time"
)

type GlobalFunction func() Person

func globalFuncImplementation() func() Person {
	return func() Person {
		return Person{
			name: "Ravi Kumar",
			born: time.Now(),
		}
	}
}

var (
	gFunc GlobalFunction = globalFuncImplementation()
)

type Person struct {
	name string
	born time.Time
}

var getPeople = func() []Person {
	var people []Person
	for i := 0; i < 3; i++ {
		people = append(people, gFunc())
	}

	return people
}()

func main() {

	times := 0

	for {

		var people []Person
		people = append(people, getPeople...)

		for i := range people {
			if times < 3 {
				people[i].born = time.Now()
			}
		}

		times += 1
		for _, p := range people {
			fmt.Println(p.born)
		}

		fmt.Println("v8 <========")

		time.Sleep(time.Second * 3)
	}
}
