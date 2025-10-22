package basics

import (
	"fmt"
)

// Doctor will be accessible to all packages but inside depends upon first letter capital or small
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func mapss() {
	statePopulations := make(map[string]int)

	statePopulations = map[string]int{
		"California": 4343,
		"Texas":      3223,
		"Florida":    3243,
		"Ohio":       324,
	}

	// m := map[[]int]string{} // invalid key type as keys should be comparable

	// Adding to map
	statePopulations["georgia"] = 32141

	// Deleting from map
	delete(statePopulations, "California")

	// checking if something is in map or not
	pop, ok := statePopulations["oho"]

	fmt.Println(pop, ok) // false if key is in map or not

	// getting length of map
	//fmt.Println(len(statePopulations))

	sp := statePopulations

	// will get deleted from both
	delete(sp, "Ohio")

	fmt.Println(sp)

	aDoctor := Doctor{
		1,
		"John Smith",
		[]string{"ma", "asd"},
	}
	fmt.Println(aDoctor)

	//*structs are independent, changing one wont change another and copies are created

	/*
		*MAPS*
			-collections of value types that are accessed via keys
			-created via literals or via make function
			-members accessed via [key] syntax
			-check for presence with value, ok form of result
			-multiple assignments refer to the same underlying data
	*/
}
