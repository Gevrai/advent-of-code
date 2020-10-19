package main

import (
	"advent-of-code-2019/utils"
	"strings"
)

func main() {
	input := utils.ReadInputFileRelative()
	_, objectMap := CreateObjectTree(input)

	count := 0
	for _, object := range objectMap {
		count += object.NumberOfParents()
	}
	println("Part 1:", count, "orbits in total")

	you := objectMap["YOU"]
	san := objectMap["SAN"]
	anc := san.FirstCommonAncestor(you)

	// number of transfers from objects YOU/SAN are oribiting to common ancestor
	youOrbit2Anc := (you.NumberOfParents() - 1) - anc.NumberOfParents()
	sanOrbit2Anc := (san.NumberOfParents() - 1) - anc.NumberOfParents()
	transfers := youOrbit2Anc + sanOrbit2Anc
	println("Part 2:", transfers, "transfers from YOU to SAN")
}

const orbitSymbol = ")"

type object struct {
	name   string
	parent *object
	orbits []*object
}

func CreateObjectTree(input []string) (root *object, objectMap map[string]*object) {
	if len(input) == 0 {
		return nil, nil
	}

	objectMap = make(map[string]*object)
	for _, line := range input {
		objects := strings.Split(line, orbitSymbol)
		if len(objects) != 2 {
			panic("invalid entry " + line)
		}

		planetName := objects[0]
		moonName := objects[1]

		planet := objectMap[planetName]
		if planet == nil {
			planet = &object{name: planetName}
			objectMap[planetName] = planet
		}

		moon := objectMap[moonName]
		if moon == nil {
			moon = &object{name: moonName}
			objectMap[moonName] = moon
		}

		moon.parent = planet
		planet.orbits = append(planet.orbits, moon)
	}

	// Return first object
	return objectMap[strings.Split(input[0], orbitSymbol)[0]], objectMap
}

func (o *object) NumberOfParents() int {
	if o == nil {
		return 0
	}
	curr := o
	i := 0
	for curr.parent != nil {
		curr = curr.parent
		i++
	}
	return i
}

func (o *object) FirstCommonAncestor(other *object) *object {
	a := o
	for a != nil {
		b := other
		for b != nil {
			if a == b {
				return a
			}
			b = b.parent
		}
		a = a.parent
	}
	return nil
}
