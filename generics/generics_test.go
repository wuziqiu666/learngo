package generics

import (
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on string", func(t *testing.T) {
		AssertEqual[string](t, "hello", "hello")
		AssertNotEqual[string](t, "hello world", "hello")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		var myStackOfInts Stack[int]
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.isEmpty())
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.isEmpty())
	})
	t.Run("string stack", func(t *testing.T) {
		var myStackOfStrings Stack[string]
		AssertTrue(t, myStackOfStrings.isEmpty())
		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.isEmpty())
		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "123")
		AssertTrue(t, myStackOfStrings.isEmpty())
	})

}
