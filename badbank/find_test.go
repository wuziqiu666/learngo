package main

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		firstEvenNumber, found := Find(number, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
	type Person struct {
		Name string
	}
	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Flower"},
			{Name: "Chris James"},
		}
		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})
		AssertTrue(t, found)
		AssertEqual(t, king.Name, "Chris James")
	})
}
