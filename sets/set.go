// set is a linear DS with no duplicate
package main

import "fmt"

type Set struct {
	integerMap map[int]bool
}

func (s *Set) NewSet() {
	s.integerMap = make(map[int]bool)
}

func (s *Set) AddElement(element int) {
	if !s.ContainsElement(element) {
		s.integerMap[element] = true
	}
}

func (s *Set) DeleteElement(element int) {
	delete(s.integerMap, element)

}

func (s *Set) ContainsElement(element int) bool {
	var exist bool
	_, exist = s.integerMap[element]
	return exist
}

func(s *Set) Intersect(another *Set) *Set{
	var intersectSet = &Set{}
	intersectSet.NewSet()
	// var value int
	for value , _ range := s.integerMap{
			if another.ContainsElement(value){
				intersectSet.AddElement(value)
			}
			
	}
	return intersectSet
}

func(s *Set)Union(anotherSet *Set) *Set{
	var unionSet = &Set{}
	unionSet.NewSet()


	for value,_ range s.integerMap{
		unionSet.AddElement(value)
	}

	for value,_ range anotherSet{
		unionSet.AddElement(value)
	}
	return unionSet
}

func main() {
	var set *Set
	set = &Set{}
	set.NewSet()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(1))
}
