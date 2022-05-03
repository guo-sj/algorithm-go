package main

import (
	"fmt"
)

type States []string

type Stations map[string]States

func (s *States) Substract(t States) {
	for _, str := range t {
		for i := range *s {
			if (*s)[i] == str {
				copy((*s)[i:], (*s)[i+1:])
				*s = (*s)[:len(*s)-1]
				break
			}
		}
	}
}

func (s *States) Intersection(t States) States {
	ret := States{}

	for _, str := range t {
		for i := range *s {
			if (*s)[i] == str {
				ret = append(ret, str)
			}
		}
	}
	return ret
}

var (
	statesNeeded    States
	currentStations Stations
	finalStations   Stations
)

func init() {
	statesNeeded = States{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	currentStations = Stations{
		"kone":   States{"id", "nv", "ut"},
		"ktwo":   States{"wa", "id", "mt"},
		"kthree": States{"or", "nv", "ca"},
		"kfour":  States{"nv", "ut"},
		"kfive":  States{"ca", "az"},
	}
	finalStations = Stations{}
}

func getResultGreedily() {
	for len(statesNeeded) > 0 {
		bestStation := ""
		bestStationCoveredStates := States{}
		for station, states := range currentStations {
			if _, ok := finalStations[station]; ok {
				continue
			}
			covered := states.Intersection(statesNeeded)
			if len(covered) > len(bestStationCoveredStates) {
				bestStation = station
				bestStationCoveredStates = covered
			}
		}
		finalStations[bestStation] = bestStationCoveredStates
		statesNeeded.Substract(bestStationCoveredStates)
	}
}

func main() {
	getResultGreedily()
	fmt.Println(finalStations)
}
