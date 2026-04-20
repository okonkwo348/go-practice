package main

import (
	"fmt"
)

func main() {

	age := map[string]int{
		"okonkwo": 23,
		"Emeka":   25,
		"Peter":   22,
	}

	// add or update
	age["okonkwo"] = 10
	age["coach"] = 50
	fmt.Println(age["okonkwo"])
	fmt.Println(age["coach"])

	// create a new map
	states := make(map[string]string)
	// var state map[string]string  panic
	states["Nigeria"] = "Lagos"
	states["Egypt"] = "Cairo"
	states["Uganda"] = "Kampala"

	fmt.Println(states)

	// to delete a pair key-value in a map.
	delete(states, "Nigeria")

	fmt.Println(states)

	state, ok := states["Uganda"]
	if ok {
		fmt.Println("it's found", state)
	} else {
		fmt.Println("it's not found", state)
	}

	for key, value := range states {
		fmt.Printf("%s is the lastest state in %s", value, key)
		fmt.Println()
	}
}
