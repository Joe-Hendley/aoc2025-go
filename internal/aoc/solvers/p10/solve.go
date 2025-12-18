package p10

import (
	"slices"
	"strings"

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
	sumFewestPresses := 0
	for _, line := range strings.Split(input, "\n") {
		machine := parseMachinePart1(line)
		result, ok := machine.searchStates()
		if !ok {
			panic("result not found!?")
		}

		sumFewestPresses += result
	}

	return sumFewestPresses
}

type machinePart1 struct {
	targetState int
	buttons     []int
	states      map[int][]bool
}

func parseMachinePart1(line string) machinePart1 {
	parts := strings.Fields(line)
	targetLightState := parseLights(parts[0])

	buttons := []int{}
	for _, part := range parts[1 : len(parts)-1] {
		buttons = append(buttons, parseButton(part))
	}

	return machinePart1{
		targetState: targetLightState,
		buttons:     buttons,
		states:      map[int][]bool{0: make([]bool, len(buttons))},
	}
}

func (m *machinePart1) searchStates() (int, bool) {
	statesToCheck := []machineStatePart1{}
	for _, button := range m.buttons {
		statesToCheck = append(statesToCheck, buildNewState(machineStatePart1{0, m.buttons}, button))
	}

	for len(statesToCheck) > 0 {
		checkedState := statesToCheck[0]

		if checkedState.state == m.targetState {
			return len(m.buttons) - len(checkedState.buttons), true
		}

		statesToCheck = statesToCheck[1:]
		for _, button := range checkedState.buttons {
			statesToCheck = append(statesToCheck, buildNewState(checkedState, button))
		}
	}

	return 0, false
}

type machineStatePart1 struct {
	state   int
	buttons []int
}

func buildNewState(ms machineStatePart1, button int) machineStatePart1 {
	newState := pushButton(ms.state, button)
	newButtons := slices.Clone(ms.buttons)
	newButtons = slices.DeleteFunc(newButtons, func(i int) bool { return i == button })

	return machineStatePart1{
		newState,
		newButtons,
	}
}

func (s *Solver) Part2(input string) int {
	return 0
}

func intPow(n, e int) int {
	p := 1

	for e > 0 {
		p *= n
		e--
	}
	return p
}

func parseLights(s string) int {
	lights := 0
	for i, c := range s[1 : len(s)-1] {
		if c == '#' {
			lights += intPow(2, i)
		}
	}

	return lights
}

func parseButton(s string) int {
	parts := must.StringSplitToInts(s[1:len(s)-1], ",")
	button := 0
	for _, part := range parts {
		button += intPow(2, part)
	}

	return button
}

func pushButton(state, button int) int {
	return state ^ button
}
