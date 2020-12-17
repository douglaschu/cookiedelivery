package main

import (
	"fmt"
	"strings" // strings package imported for edits to printed output
)

func main() {
	fmt.Printf("Doll Delivery\n")

	// load sample neighborhood from data in build requirements
	neighborhood := newGraph()
	neighborhood.addStreet("Natalie's abode", "Laura's crib", 9)
	neighborhood.addStreet("Natalie's abode", "Greg's casa", 4)
	neighborhood.addStreet("Natalie's abode", "Chloe's pad", 18)
	neighborhood.addStreet("Natalie's abode", "Sophia's apartment", 8)
	neighborhood.addStreet("Sophia's apartment", "Jordan's condo", 7)
	neighborhood.addStreet("Sophia's apartment", "Ian's dwelling", 17)
	neighborhood.addStreet("Greg's casa", "Ian's dwelling", 13)
	neighborhood.addStreet("Greg's casa", "Ben's digs", 19)
	neighborhood.addStreet("Greg's casa", "Chloe's pad", 14)
	neighborhood.addStreet("Jordan's condo", "Keegan's farm", 10)
	neighborhood.addStreet("Jordan's condo", "Nathan's flat", 11)
	neighborhood.addStreet("Jordan's condo", "Danielle's den", 6)
	neighborhood.addStreet("Chloe's pad", "Laura's crib", 19)
	neighborhood.addStreet("Chloe's pad", "Nathan's flat", 15)
	neighborhood.addStreet("Chloe's pad", "Jake's haunt", 14)
	neighborhood.addStreet("Laura's crib", "Keegan's farm", 9)
	neighborhood.addStreet("Laura's crib", "Nathan's flat", 12)
	neighborhood.addStreet("Danielle's den", "Jake's haunt", 10)
	neighborhood.addStreet("Danielle's den", "Ben's digs", 9)
	neighborhood.addStreet("Ben's digs", "Ian's dwelling", 20)
	neighborhood.addStreet("Ben's digs", "Nathan's flat", 12)
	neighborhood.addStreet("Ian's dwelling", "Jake's haunt", 18)
	neighborhood.addStreet("Nathan's flat", "Keegan's farm", 3)

	// Test 1: Check pathfinder function vs. example answer in build reqs for Natalie's to Jake's.
	fmt.Println("\nTest 1: Natalie's to Jake's")
	// formatting for a readable output that matches build req.
	dist, path := neighborhood.getPath("Natalie's abode", "Jake's haunt")
	// format verbs used to manipulate getPath() output to add labels on printing
	fmt.Printf("Distance: %d\nPath: [%s]", dist, strings.Join(path, ", "))
	// strings.Join lets me put a comma between the street sequence in the path output for readability.

	// Test 2: Check if two houses are directly adjacent: Natalie's and Laura's.
	fmt.Println("\nTest 2: Natalie's to Chloe's")
	dist2, path2 := neighborhood.getPath("Natalie's abode", "Chloe's pad")
	fmt.Printf("Distance: %d\nPath: [%s]", dist2, strings.Join(path2, ", "))

	// Test 3: Check if function works in reverse (Jake's to Natalie's)
	fmt.Println("\nTest 3: Jake's to Natalie's")
	dist3, path3 := neighborhood.getPath("Jake's haunt", "Natalie's abode")
	fmt.Printf("Distance: %d\nPath: [%s]", dist3, strings.Join(path3, ", "))

	// Test 4: Let's try another route: Ian's to Chloe's.
	fmt.Println("\nTest 4: Ian's to Chloe's")
	dist4, path4 := neighborhood.getPath("Ian's dwelling", "Chloe's pad")
	fmt.Printf("Distance: %d\nPath: [%s]", dist4, strings.Join(path4, ", "))

	// Test 5: What if someone double-enters a house as both start and target location?
	fmt.Println("\nTest 5: Natalie's to Natalie's")
	dist5, path5 := neighborhood.getPath("Natalie's abode", "Natalie's abode")
	fmt.Printf("Distance: %d\nPath: [%s]", dist5, strings.Join(path5, ", "))

}
