package advanced

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (ps *personSorter) Len() int {
	return len(ps.people)
}

func (ps *personSorter) Less(i, j int) bool {
	return ps.by(&ps.people[i], &ps.people[j])
}

func (ps *personSorter) Swap(i, j int) {
	ps.people[i], ps.people[j] = ps.people[j], ps.people[i]
}

func (by By) Sort(people []Person) {
	p := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(p)
}

type ByAge []Person
type ByName []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func ExampleSorting() {

	// Built-in functions
	numbers := []int{4, 2, 1, 3}
	sort.Ints(numbers)

	stringSlice := []string{"c", "a", "bas"}
	sort.Strings(stringSlice)

	// Sorting By Functions
	people := []Person{
		{"Alice", 32},
		{"Bob", 25},
		{"Charlie", 30},
	}

	fmt.Println("Original:", people)

	sort.Sort(ByAge(people))
	fmt.Println("Sorted by age:", people)

	sort.Sort(ByName(people))
	fmt.Println("Sorted by name:", people)

	// ==================================================================
	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(age).Sort(people)
	fmt.Println("Sorted by age:", people)

	// =================================================================
	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	By(name).Sort(people)
	fmt.Println("Sorted by name:", people)
}

/*
# Built in Functions
	- sort.Ints()
	- sort.Sort(sort.Interface)
	- sort.Strings

# Sorting By Functions - (Custom Criteria, Reusability, Flexibility)
	- `sort.Interface` consists of three methods
		- Len() int
		- Less(i, j int) bool
		- Swap(i, j int)


*/
