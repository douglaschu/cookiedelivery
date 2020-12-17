package main

import (
	"fmt"
	"strings" // strings package imported for edits to printed output
)

func main() {
	fmt.Printf("Doll Delivery\n")

	// load sample neighborhood from data in build requirements
	neighborhood := newGraph()
	neighborhood.addStreet("Kruthika's abode", "Mark's crib", 9)
	neighborhood.addStreet("Kruthika's abode", "Greg's casa", 4)
	neighborhood.addStreet("Kruthika's abode", "Matt's pad", 18)
	neighborhood.addStreet("Kruthika's abode", "Brian's apartment", 8)
	neighborhood.addStreet("Brian's apartment", "Wesley's condo", 7)
	neighborhood.addStreet("Brian's apartment", "Cam's dwelling", 17)
	neighborhood.addStreet("Greg's casa", "Cam's dwelling", 13)
	neighborhood.addStreet("Greg's casa", "Mike's digs", 19)
	neighborhood.addStreet("Greg's casa", "Matt's pad", 14)
	neighborhood.addStreet("Wesley's condo", "Kirk's farm", 10)
	neighborhood.addStreet("Wesley's condo", "Nathan's flat", 11)
	neighborhood.addStreet("Wesley's condo", "Bryce's den", 6)
	neighborhood.addStreet("Matt's pad", "Mark's crib", 19)
	neighborhood.addStreet("Matt's pad", "Nathan's flat", 15)
	neighborhood.addStreet("Matt's pad", "Craig's haunt", 14)
	neighborhood.addStreet("Mark's crib", "Kirk's farm", 9)
	neighborhood.addStreet("Mark's crib", "Nathan's flat", 12)
	neighborhood.addStreet("Bryce's den", "Craig's haunt", 10)
	neighborhood.addStreet("Bryce's den", "Mike's digs", 9)
	neighborhood.addStreet("Mike's digs", "Cam's dwelling", 20)
	neighborhood.addStreet("Mike's digs", "Nathan's flat", 12)
	neighborhood.addStreet("Cam's dwelling", "Craig's haunt", 18)
	neighborhood.addStreet("Nathan's flat", "Kirk's farm", 3)

	// Test 1: Check pathfinder function vs. example answer in build reqs for Kruthika's to Craig's.
	fmt.Println("\nTest 1: Kruthika's to Craig's")
	// formatting for a readable output that matches build req.
	dist, path := neighborhood.getPath("Kruthika's abode", "Craig's haunt")
	// format verbs used to manipulate getPath() output to add labels on printing
	fmt.Printf("Distance: %d\nPath: [%s]", dist, strings.Join(path, ", "))
	// strings.Join lets me put a comma between the street sequence in the path output for readability.

	// Test 2: Check if two houses are directly adjacent: Kruthika's and Mark's.
	fmt.Println("\nTest 2: Kruthika's to Matt's")
	dist2, path2 := neighborhood.getPath("Kruthika's abode", "Matt's pad")
	fmt.Printf("Distance: %d\nPath: [%s]", dist2, strings.Join(path2, ", "))

	// Test 3: Check if function works in reverse (Craig's to Kruthika's)
	fmt.Println("\nTest 3: Craig's to Kruthika's")
	dist3, path3 := neighborhood.getPath("Craig's haunt", "Kruthika's abode")
	fmt.Printf("Distance: %d\nPath: [%s]", dist3, strings.Join(path3, ", "))

	// Test 4: Let's try another route: Cam's to Matt's.
	fmt.Println("\nTest 4: Cam's to Matt's")
	dist4, path4 := neighborhood.getPath("Cam's dwelling", "Matt's pad")
	fmt.Printf("Distance: %d\nPath: [%s]", dist4, strings.Join(path4, ", "))

	// Test 5: What if someone double-enters a house as both start and target location?
	fmt.Println("\nTest 5: Kruthika's to Kruthika's")
	dist5, path5 := neighborhood.getPath("Kruthika's abode", "Kruthika's abode")
	fmt.Printf("Distance: %d\nPath: [%s]", dist5, strings.Join(path5, ", "))

}
