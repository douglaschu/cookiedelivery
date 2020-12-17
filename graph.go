package main

// typing for street properties
type street struct {
	house    string
	distance int
}

// typing for graph properties
type graph struct {
	houses map[string][]street
}

// creates a map based on the streets
func newGraph() *graph {
	return &graph{houses: make(map[string][]street)}
}

// this function allows for building out the graph by adding streets
// input format is (house A), (house B), (the distance between A and B)
func (g *graph) addStreet(startLocation, targetLocation string, distance int) {
	g.houses[startLocation] = append(g.houses[startLocation], street{house: targetLocation, distance: distance})
	g.houses[targetLocation] = append(g.houses[targetLocation], street{house: startLocation, distance: distance})
}

// returns path as described as a sequences of houses visited
func (g *graph) getStreets(house string) []street {
	return g.houses[house]
}

// This function calculates the distance of the shortest path between two houses (houses) as well as the path itself, described in order of houses visited.
func (g *graph) getPath(startLocation, targetLocation string) (int, []string) {
	h := newHeap()
	h.push(path{value: 0, houses: []string{startLocation}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest unvisited house
		p := h.pop()
		house := p.houses[len(p.houses)-1]

		// if house is visited, continue to next house
		if visited[house] {
			continue
		}

		//stops and returns result if nearest house is targetLocation
		if house == targetLocation {
			return p.value, p.houses
		}

		for _, e := range g.getStreets(house) {
			if !visited[e.house] {
				// This adds the cost of moving to this house to the total cost accrued so far on the path.
				h.push(path{value: p.value + e.distance, houses: append([]string{}, append(p.houses, e.house)...)})
			}
		}
		//sets visited property of house to true
		visited[house] = true

	}
	// returns a value of zero if there's no distance between houses
	// (As in: you entered the same house for both start and target.)
	return 0, nil

}
