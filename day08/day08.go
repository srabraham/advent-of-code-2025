package day08

import (
	_ "embed"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input0.txt
var input0 string

//go:embed input1.txt
var input1 string

type XYZ struct {
	X, Y, Z int
}

type XYZPairDist struct {
	A, B XYZ
	Dist int
}

func distSq(a, b XYZ) int {
	return (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y) + (a.Z-b.Z)*(a.Z-b.Z)
}

func Part1(input string, numConnections int) int {
	// list of all junctions (vertices)
	var junctions []XYZ
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		junctions = append(junctions, mustXYZ(line))
	}

	// a junction pair and the distance-squared between them
	var dists []XYZPairDist
	for i, a := range junctions {
		for j, b := range junctions {
			if j >= i {
				break
			}
			dists = append(dists, XYZPairDist{
				A:    a,
				B:    b,
				Dist: distSq(junctions[i], junctions[j]),
			})
		}
	}

	// sort all the pairs by the distances ascending
	slices.SortFunc(dists, func(a, b XYZPairDist) int {
		return a.Dist - b.Dist
	})

	// make size-1 circuits out of every junction
	type groupInd int
	groups := make(map[groupInd][]XYZ)
	for i, j := range junctions {
		groups[groupInd(i)] = []XYZ{j}
	}

	// connect the circuits together
	for i := range numConnections {
		// find the groups for each junction
		var groupA, groupB groupInd
		for kA, vA := range groups {
			if slices.Contains(vA, dists[i].A) {
				groupA = kA
			}
		}
		for kB, vB := range groups {
			if slices.Contains(vB, dists[i].B) {
				groupB = kB
			}
		}
		if groupA == groupB {
			// already in the same group, nothing to do
			continue
		}
		groups[groupA] = append(groups[groupA], groups[groupB]...)
		groups[groupB] = nil
	}

	var circuitSizes []int
	for _, v := range groups {
		circuitSizes = append(circuitSizes, len(v))
	}
	slices.Sort(circuitSizes)
	slices.Reverse(circuitSizes)

	return circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
}

func Part2(input string) int {
	// list of all junctions (vertices)
	var junctions []XYZ
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		junctions = append(junctions, mustXYZ(line))
	}

	// a junction pair and the distance-squared between them
	var dists []XYZPairDist
	for i, a := range junctions {
		for j, b := range junctions {
			if j >= i {
				break
			}
			dists = append(dists, XYZPairDist{
				A:    a,
				B:    b,
				Dist: distSq(junctions[i], junctions[j]),
			})
		}
	}

	// sort all the pairs by the distances ascending
	slices.SortFunc(dists, func(a, b XYZPairDist) int {
		return a.Dist - b.Dist
	})

	// make size-1 circuits out of every junction
	type groupInd int
	groups := make(map[groupInd][]XYZ)
	for i, j := range junctions {
		groups[groupInd(i)] = []XYZ{j}
	}

	// connect the circuits together
	var i int
	for i = 0; len(groups) > 1; i++ {
		// find the groups for each junction
		var groupA, groupB groupInd
		for kA, vA := range groups {
			if slices.Contains(vA, dists[i].A) {
				groupA = kA
			}
		}
		for kB, vB := range groups {
			if slices.Contains(vB, dists[i].B) {
				groupB = kB
			}
		}
		//log.Println("connecting", dists[i], "len(groups) =", len(groups))
		if groupA == groupB {
			// already in the same group, nothing to do
			continue
		}
		groups[groupA] = append(groups[groupA], groups[groupB]...)
		delete(groups, groupB)
	}

	//log.Println("last connected junctions were", dists[i-1])
	return dists[i-1].A.X * dists[i-1].B.X
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mustXYZ(s string) XYZ {
	splits := strings.Split(s, ",")
	if len(splits) != 3 {
		log.Panicln("wrong input length", s)
	}
	return XYZ{
		X: mustInt(splits[0]),
		Y: mustInt(splits[1]),
		Z: mustInt(splits[2]),
	}
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	must(err)
	return i
}
