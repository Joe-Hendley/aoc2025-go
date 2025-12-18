package p11

import (
	"slices"
	"strings"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/fun"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/logger"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func (s *Solver) Part1(input string) int {
	devices := populateDevices(input)

	return visit(devices, devices["you"], []string{})
}

func (s *Solver) Part2(input string) int {
	//devices := populateDevices(input)

	return 0
	//return visitComplex(devices, devices["svr"], []string{}, false, false)
}

func populateDevices(input string) map[string]device {
	deviceList := fun.Map(strings.Split(input, "\n"), parseDevice)
	devices := make(map[string]device, len(deviceList))
	fun.Apply(deviceList, func(d device) {
		if existingDevice, ok := devices[d.id]; ok {
			d.inputs = existingDevice.inputs
		}
		devices[d.id] = d

		for _, outputID := range d.outputs {
			if existingDevice, ok := devices[outputID]; ok {
				existingDevice.inputs = append(existingDevice.inputs, d.id)
				devices[outputID] = existingDevice
			} else {
				devices[outputID] = device{
					id:     outputID,
					inputs: []string{d.id},
				}
			}
		}
	})

	return devices
}

type device struct {
	id      string
	inputs  []string
	outputs []string
}

func parseDevice(line string) device {
	parts := strings.Split(line, ":")
	return device{
		id:      parts[0],
		outputs: strings.Fields(parts[1]),
	}
}

func visit(devices map[string]device, d device, visited []string) int {
	if slices.Contains(visited, d.id) {
		return 0
	}

	if d.id == "out" {
		return 1
	}

	s := 0
	for _, newDeviceID := range d.outputs {
		newVisited := slices.Clone(visited)
		newVisited = append(newVisited, d.id)
		s += visit(devices, devices[newDeviceID], newVisited)
	}
	return s
}

func visitComplex(devices map[string]device, d device, visited []string, visitedFFT, visitedDAC bool) int {
	if slices.Contains(visited, d.id) {
		return 0
	}

	if d.id == "fft" {
		visitedFFT = true
	}

	if d.id == "dac" {
		visitedDAC = true
	}

	if d.id == "out" {
		if visitedFFT && visitedDAC {
			return 1
		}

		return 0
	}

	s := 0
	for _, newDeviceID := range d.outputs {
		newVisited := slices.Clone(visited)
		newVisited = append(newVisited, d.id)
		s += visitComplex(devices, devices[newDeviceID], newVisited, visitedFFT, visitedDAC)
	}
	return s
}
