package p8

import (
	"math"
	"slices"
	"strings"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/grid"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/logger"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/must"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func (s *Solver) Part1(input string) int {
	// test data wants 10 connections, real data wants 1000
	// test data has 20 lines, real has 1000 so let's use an ugly if statement

	numConnections := 10
	if len(input) > 1000 {
		numConnections = 1000
	}

	playground := newPlayground(input, s.Logger)

	// make the n shortest connections
	connections := 0
	for connections < numConnections {
		closestPair, ok := playground.popClosestPair()
		if !ok {
			panic("out of possible connections")
		}

		playground.connect(closestPair.junctionBoxIDs.a, closestPair.junctionBoxIDs.b)

		// WHY WOULD YOU CONNECT THEM IF THEY'RE ALREADY PART OF THE SAME CIRCUIT I LOST HOURS OF MY LIFE ON THIS
		// MAKE IT CLEARER IF YOU SAY NOTHING HAPPENS FFS
		connections++
	}

	// then start counting circuit sizes

	circuitSizes := []int{}
	for _, circuit := range playground.circuits {
		circuitSizes = append(circuitSizes, len(circuit))
	}

	slices.Sort(circuitSizes)
	slices.Reverse(circuitSizes)

	return circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
}

func (s *Solver) Part2(input string) int {
	playground := newPlayground(input, s.Logger)
	lastBoxes := pair{}

	for len(playground.circuits) > 1 {
		closestPair, ok := playground.popClosestPair()
		if !ok {
			panic("out of possible connections")
		}
		lastBoxes = closestPair.junctionBoxIDs
		playground.connect(closestPair.junctionBoxIDs.a, closestPair.junctionBoxIDs.b)
	}

	// then start counting circuit sizes

	x1 := playground.junctionBoxes[lastBoxes.a].x
	x2 := playground.junctionBoxes[lastBoxes.b].x

	return x1 * x2
}

func euclidianDistance(x1, x2, y1, y2, z1, z2 float64) float64 {
	xComponent := math.Pow(x1-x2, 2)
	yComponent := math.Pow(y1-y2, 2)
	zComponent := math.Pow(z1-z2, 2)
	return math.Sqrt(xComponent + yComponent + zComponent)
}

type junctionBox struct {
	id        int
	x, y, z   int
	circuitID int
}

func (jb junctionBox) distance(jb2 junctionBox) float64 {
	return euclidianDistance(
		float64(jb.x), float64(jb2.x),
		float64(jb.y), float64(jb2.y),
		float64(jb.z), float64(jb2.z),
	)
}

func parseLine(input string, id int) junctionBox {
	intCoords := must.StringSplitToInts(input, ",")
	return junctionBox{
		id:        id,
		x:         intCoords[0],
		y:         intCoords[1],
		z:         intCoords[2],
		circuitID: id,
	}
}

type pair struct {
	a, b int
}

type distance struct {
	junctionBoxIDs pair
	value          float64
}

func distanceSortFunc(a, b distance) int {
	switch {
	case a.value < b.value:
		return -1
	case b.value < a.value:
		return 1
	default:
		return 0
	}
}

type playground struct {
	logger        logger.Logger
	junctionBoxes []junctionBox
	distances     grid.Grid[distance]
	closestPairs  []distance
	circuits      map[int][]int
}

func newPlayground(input string, logger logger.Logger) playground {
	p := playground{logger: logger}
	p.parseJunctionBoxes(input)
	return p
}

func (p *playground) parseJunctionBoxes(input string) {
	lines := strings.Split(input, "\n")
	junctionBoxes := make([]junctionBox, len(lines))
	for i, line := range lines {
		junctionBoxes[i] = parseLine(line, i)
	}

	circuits := map[int][]int{}
	for i := range junctionBoxes {
		circuits[i] = []int{i}
	}

	distances := grid.New(len(junctionBoxes), len(junctionBoxes), distance{})
	closestPairs := []distance{}

	for boxIndex := range junctionBoxes {
		for targetBoxIndex := range junctionBoxes {
			if targetBoxIndex <= boxIndex {
				continue
			}

			d := distance{
				junctionBoxIDs: pair{
					a: boxIndex,
					b: targetBoxIndex,
				},
				value: junctionBoxes[boxIndex].distance(junctionBoxes[targetBoxIndex]),
			}
			distances.Replace(boxIndex, targetBoxIndex, d)
			closestPairs = append(closestPairs, d)
		}
	}

	slices.SortFunc(closestPairs, distanceSortFunc)

	p.junctionBoxes = junctionBoxes
	p.circuits = circuits
	p.distances = distances
	p.closestPairs = closestPairs
}

func (p *playground) popClosestPair() (distance, bool) {
	if len(p.closestPairs) < 1 {
		return distance{}, false
	}

	closest := p.closestPairs[0]
	p.closestPairs = p.closestPairs[1:]

	return closest, true
}

func (p *playground) connect(junctionBox1, junctionBox2 int) {
	circuitID1 := p.junctionBoxes[junctionBox1].circuitID
	circuitID2 := p.junctionBoxes[junctionBox2].circuitID

	if circuitID1 == circuitID2 {
		return
	}

	p.joinCircuits(circuitID1, circuitID2)

	for _, boxID := range p.circuits[circuitID1] {
		p.junctionBoxes[boxID].circuitID = circuitID1
	}

	return
}

func (p *playground) joinCircuits(circuitID1, circuitID2 int) {
	joined := slices.Concat(p.circuits[circuitID1], p.circuits[circuitID2])
	p.circuits[circuitID1] = joined
	delete(p.circuits, circuitID2)
}
