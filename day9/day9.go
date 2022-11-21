package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type connection struct {
	city1    string
	city2    string
	distance int
}

func (c *connection) containsBoth(cityA, cityB string) bool {
	if cityA == c.city1 && cityB == c.city2 {
		return true
	}
	return false
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\r\n")

	connections := make([]connection, 0)
	cities := make([]string, 0)
	for _, l := range lines {
		splitted := strings.Split(l, " ")
		distance, err := strconv.ParseInt(splitted[len(splitted)-1], 10, 64)
		if err != nil {
			panic(err)
		}
		connections = append(connections, connection{city1: splitted[0], city2: splitted[2], distance: int(distance)})
		connections = append(connections, connection{city1: splitted[2], city2: splitted[0], distance: int(distance)})
		addToCities(&cities, splitted[0], splitted[2])
	}

	routes := findAllRoutes("", 0, cities, connections)
	shortest := -1
	for _, r := range routes {
		if shortest == -1 {
			shortest = r
		} else {
			if shortest > r {
				shortest = r
			}
		}
	}
	fmt.Printf("Day 9 - part 1\n")
	fmt.Printf("The shortest route is '%d' long\n", shortest)

	longest := -1
	for _, r := range routes {
		if shortest == -1 {
			longest = r
		} else {
			if longest < r {
				longest = r
			}
		}
	}
	fmt.Printf("Day 9 - part 2\n")
	fmt.Printf("The shortest route is '%d' long\n", longest)
}

func addToCities(cities *[]string, toAdd ...string) {
	for _, s := range toAdd {
		if !containsCity(cities, s) {
			*cities = append(*cities, s)
		}
	}
}

func containsCity(cities *[]string, city string) bool {
	for _, c := range *cities {
		if c == city {
			return true
		}
	}
	return false
}

func findAllRoutes(prevCity string, currentDistance int, cities []string, connections []connection) []int {
	if len(cities) == 1 && prevCity != "" {
		way := getCurrentConnection(prevCity, cities[0], connections) + currentDistance
		return []int{way}
	}

	result := make([]int, 0)
	for i, c := range cities {
		distance := getCurrentConnection(prevCity, c, connections) + currentDistance

		copyCities := make([]string, len(cities))
		copy(copyCities, cities)
		remainingCities := append(copyCities[:i], copyCities[i+1:]...)
		result = append(result, findAllRoutes(c, distance, remainingCities, connections)...)
	}
	return result
}

func getCurrentConnection(cityA, cityB string, connections []connection) int {
	for _, c := range connections {
		if c.containsBoth(cityA, cityB) {
			return c.distance
		}
	}
	return 0
}
