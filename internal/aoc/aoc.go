package aoc

import (
	"flag"
	"fmt"
	"time"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/file"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/solvers/p1"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/solvers/p2"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/solvers/p3"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/solvers/p4"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/solvers/p5"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/solvers/p6"
)

const header = `
  ░█████╗░░█████╗░░█████╗░░░██████╗░░█████╗░██████╗░███████╗
  ██╔══██╗██╔══██╗██╔══██╗░░╚════██╗██╔══██╗╚════██╗██╔════╝
  ███████║██║░░██║██║░░╚═╝░░░░███╔═╝██║░░██║░░███╔═╝███████╗
  ██╔══██║██║░░██║██║░░██╗░░██╔══╝░░██║░░██║██╔══╝░░╚════██║
  ██║░░██║╚█████╔╝╚█████╔╝░░███████╗╚█████╔╝███████╗███████║
  ╚═╝░░╚═╝░╚════╝░░╚════╝░░░╚══════╝░╚════╝░╚══════╝╚══════╝`

type solver interface {
	Init(verbose bool)
	Part1(input string) int
	Part2(input string) int
}

var solvers = []solver{
	&p1.Solver{},
	&p2.Solver{},
	&p3.Solver{},
	&p4.Solver{},
	&p5.Solver{},
	&p6.Solver{},
}

func Solve() int {
	fmt.Println(header)
	fmt.Println("")

	cfg := newConfigFromFlags()

	if cfg.puzzle == 0 {
		return solveAll(cfg)
	} else {
		return solve(cfg.puzzle, cfg.verbose, cfg.bench)
	}
}

// TODO: not global
var totalTime time.Duration
var totalStars = 0

func solveAll(cfg config) int {

	pooledOut := 0

	for i := range solvers {
		out := solve(i+1, cfg.verbose, cfg.bench)
		if out != 0 {
			pooledOut = -1
		}

		fmt.Println("")
	}

	if cfg.bench {
		fmt.Printf("%33s⏱️ %-8s\n", "", buildTimeString(totalTime))
	}

	fmt.Println()

	fmt.Println(buildStars(totalStars))

	return pooledOut
}

func solve(puzzle int, verbose, bench bool) int {
	if len(solvers) < puzzle || puzzle < 1 {
		fmt.Printf("invalid puzzle #%d, valid range is 1 to %d\n", puzzle, len(solvers))
		return -1
	}

	path := fmt.Sprintf("./input/%d.txt", puzzle)

	input, err := file.ReadToString(path)
	if err != nil {
		fmt.Printf("could not find input for puzzle %d at %s\n", puzzle, path)
		return -1
	}

	solver := solvers[puzzle-1]

	solver.Init(verbose)

	now := time.Now()
	result1 := solver.Part1(input)
	elapsed := time.Since(now)

	benchStr := "\t"

	if bench {
		totalTime += elapsed
		benchStr = fmt.Sprintf("⏱️ %-8s\t", buildTimeString(elapsed))
	}

	if result1 != 0 {
		totalStars += 1
		fmt.Printf("\t📅%2d p1: %-16d%s%s\n", puzzle, result1, benchStr, decorate(puzzle, 1))
	}

	now = time.Now()
	result2 := solver.Part2(input)
	elapsed = time.Since(now)

	benchStr = "\t"

	if bench {
		totalTime += elapsed
		benchStr = fmt.Sprintf("⏱️ %-8s\t", buildTimeString(elapsed))
	}

	if result2 != 0 {
		totalStars += 1
		fmt.Printf("\t     p2: %-16d%s%8s\n", result2, benchStr, decorate(puzzle, 2))
	}

	return 0
}

type config struct {
	verbose bool
	puzzle  int
	bench   bool
}

func newConfigFromFlags() config {
	verbose := flag.Bool("v", false, "enable verbose logging")
	puzzle := flag.Int("p", 0, "specific puzzle to solve")
	bench := flag.Bool("bench", false, "enable and display benchmarks")

	flag.Parse()

	return config{
		verbose: *verbose,
		puzzle:  *puzzle,
		bench:   *bench,
	}
}
