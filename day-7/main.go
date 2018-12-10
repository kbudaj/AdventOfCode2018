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

func mapToOrderedKeys(m map[int]*step) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}

func partOne(inp []string) []string {
	stepMap := parseInput(inp)
	orderedStepIDs := mapToOrderedKeys(stepMap)
	desiredLength := len(orderedStepIDs)
	finished := make([]string, 0, len(orderedStepIDs))

	i := 0
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
	return finished
}

type worker struct {
	timeLeft int
	task     *step
}

func getAvailableStep(stepMap map[int]*step, orderedStepIDs *[]int) *step {
	cycle := len(*orderedStepIDs)
	for i := 0; i < cycle; i++ {
		stepID := (*orderedStepIDs)[i]
		if !stepMap[stepID].isBlocked(&stepMap) {
			newOrdStepIDs := *orderedStepIDs
			newOrdStepIDs = append(newOrdStepIDs[:i], newOrdStepIDs[i+1:]...)
			*orderedStepIDs = newOrdStepIDs
			return stepMap[stepID]
		}
	}
	return nil
}

func partTwo(inp []string) int {
	stepMap := parseInput(inp)
	k := mapToOrderedKeys(stepMap)
	orderedStepIDs := &k
	w1, w2, w3, w4, w5 := worker{0, nil}, worker{0, nil}, worker{0, nil}, worker{0, nil}, worker{0, nil}
	workers := []*worker{&w1, &w2, &w3, &w4, &w5}
	var seconds int
	for seconds = -1; len(*orderedStepIDs) > 0; seconds++ {
		for _, w := range workers {
			if w.timeLeft > 0 {
				w.timeLeft--
			}
			if w.timeLeft == 0 {
				if w.task != nil {
					w.task.finished = true
					w.task = nil
				}
				task := getAvailableStep(stepMap, orderedStepIDs)
				if task != nil {
					w.task = task
					w.timeLeft = 60 + task.id - 64 // ASCII A 65-64=1
				}
			}
		}
	}
	seconds--
	for _, w := range workers {
		seconds += w.timeLeft
	}
	return seconds
}

func main() {
	inp := make([]string, 0, 256)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := s.Text()
		inp = append(inp, v)
	}
	fmt.Println("1: ", strings.Join(partOne(inp), ""))
	fmt.Println("2: ", partTwo(inp))
}
