package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type step struct {
	id         int
	finished   bool
	conditions *[]*step
}

func (s *step) isBlocked(stepMap *map[int]*step) bool {
	for _, c := range *s.conditions {
		if !c.finished {
			return true
		}
	}
	return false
}

func createStep(stepID int) step {
	conditions := make([]*step, 0, 8)
	p := step{stepID, false, &conditions}
	return p
}

func parseInput(inp []string) map[int]*step {
	re := regexp.MustCompile(`Step ([A-Z]).+([A-Z])`)
	stepMap := make(map[int]*step)
	for _, v := range inp {
		inputArr := re.FindStringSubmatch(v)
		conditionID := int(inputArr[1][0])
		stepID := int(inputArr[2][0])

		var condition *step
		if stepMap[conditionID] == nil {
			c := createStep(conditionID)
			condition = &c
			stepMap[conditionID] = condition
		} else {
			condition = stepMap[conditionID]
		}

		var currentStep *step
		if stepMap[stepID] == nil {
			s := createStep(stepID)
			currentStep = &s
			stepMap[stepID] = currentStep
		} else {
			currentStep = stepMap[stepID]
		}

		conditions := append(*currentStep.conditions, condition)
		*currentStep.conditions = conditions
	}
	return stepMap
}

func partOne(inp []string) string {
	stepMap := parseInput(inp)

	// Map keys to ordered list
	orderedStepIDs := make([]int, len(stepMap))
	i := 0
	for k := range stepMap {
		orderedStepIDs[i] = k
		i++
	}
	sort.Ints(orderedStepIDs)

	desiredLength := len(orderedStepIDs)
	finished := make([]string, 0, len(orderedStepIDs))

	i = 0
	for len(finished) < desiredLength {
		stepID := orderedStepIDs[i]
		if !stepMap[stepID].isBlocked(&stepMap) {
			stepMap[stepID].finished = true
			finished = append(finished, string(stepMap[stepID].id))
			orderedStepIDs = append(orderedStepIDs[:i], orderedStepIDs[i+1:]...)
			i = 0
		} else {
			i++
		}
	}
	return strings.Join(finished, "")
}

func main() {
	inp := make([]string, 0, 256)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := s.Text()
		inp = append(inp, v)
	}
	fmt.Println("1: ", partOne(inp))
	// fmt.Println("2: ", partTwo(points))
}
