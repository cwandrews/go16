package main

import (
	"fmt"
	"sort"
)

type People struct{ first, last string }

func (p People) String() string { return p.first + " " + p.last }

type ByFirstName []People

// Here we implement the interface described by the sort.Interface type for
// ByFirstName, making our custom type sortable. https://golang.org/pkg/sort/#Interface
func (s ByFirstName) Len() int           { return len(s) }
func (s ByFirstName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByFirstName) Less(i, j int) bool { return s[i].first < s[j].first }

func sort_some_people() []People {
	people := []People{
		{"Bill", "Clinton"}, {"Romeo", "Capulet"},
		{"Bill", "Cosby"}, {"Juliet", "Montague"},
		{"Bill", "Gates"}, {"Walter", "White"},
		{"April", "O'Neal"},
	}
	sort.Sort(ByFirstName(people))
	return people
}

func main() {
	peeps := sort_some_people()
	fmt.Println(peeps)

}
