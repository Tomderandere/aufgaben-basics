package check_ordering

import "fmt"

func ExampleCheckOrdering() {
	fmt.Println(CheckOrdering([]string{"a", "b", "c", "d", "e", "f"}, "a", "f"))
	fmt.Println(CheckOrdering([]string{"a", "b", "c", "d", "e", "f"}, "f", "a"))
	fmt.Println(CheckOrdering([]string{"a", "b", "c", "e", "d", "f"}, "a", "f"))
	fmt.Println(CheckOrdering([]string{"b", "a", "c", "e", "f", "d"}, "a", "f"))
	fmt.Println(CheckOrdering([]string{"a", "b", "c", "d", "e", "f"}, "f", "g"))
	fmt.Println(CheckOrdering([]string{}, "a", "b"))

	// Output:
	// true
	// false
	// true
	// true
	// false
	// false
}
