package must

import (
	"fmt"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func StringFieldsToInts(s string) []int {
	fields := strings.Fields(s)
	ints := make([]int, len(fields))

	for i := range fields {
		ints[i] = Atoi(fields[i])
	}

	return ints
}

func StringToLinesOfInts(s string) [][]int {
	lines := strings.Split(s, "\n")
	intLines := [][]int{}

	for _, line := range lines {
		intLines = append(intLines, StringFieldsToInts(line))
	}

	return intLines
}

func StringSplitToInts(s string, sep string) []int {
	split := strings.Split(s, sep)
	ints := make([]int, len(split))

	for i := range split {
		var err error
		ints[i], err = strconv.Atoi(split[i])
		if err != nil {

			panic(fmt.Sprintf("parsing {%s} delim {%s} err {%v}", s, sep, err))
		}
	}

	return ints
}
