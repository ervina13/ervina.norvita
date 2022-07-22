package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	ValueA := s[i].Correct*4 - s[i].Wrong
	ValueB := s[j].Correct*4 - s[j].Wrong

	//if value ==
	if ValueA == ValueB {
		//correct ==
		if s[i].Correct == s[j].Correct {
			//nama
			if s[i].Name > s[j].Name {
				return true
			}
			//correct i < j
		} else if s[i].Correct < s[j].Correct {
			return true
		}
		//value A < B
	} else if ValueA < ValueB {
		return true
	}

	return false // TODO: replace this
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1; j++ {
			isLess := s.Less(j, j+1)
			if isLess {
				s.Swap(j, j+1)
			}
		}
	}
	//sort.Sort(s)
	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},  //10 |3 |-12
		{"Agus", 3, 4, 0},  //8  |4 |-12
		{"Doni", 3, 0, 7},  //12 |1 |-12
		{"Ega", 3, 0, 7},   //12 |2 |-12
		{"Anton", 2, 0, 5}, //8  |5 |- 8
	})
	sort.Sort(scores)
	fmt.Println(scores)
	fmt.Println(scores.TopStudents())
}
