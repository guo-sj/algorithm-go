package main

import (
	"fmt"
)

var (
	netGraph          map[string][]string
	personQueen       []string
	personStack       []string
	bfsSearchedPerson map[string]bool
	dfsSearchedPerson map[string]bool
)

func init() {
	netGraph = make(map[string][]string)
	netGraph["you"] = []string{"alice", "bob", "claire"}
	netGraph["bob"] = []string{"anuj", "peggy"}
	netGraph["alice"] = []string{"peggy"}
	netGraph["claire"] = []string{"thom", "jonny"}
	netGraph["anuj"] = []string{}
	netGraph["peggy"] = []string{}
	netGraph["thom"] = []string{}
	netGraph["jonny"] = []string{}

	bfsSearchedPerson = make(map[string]bool)
	dfsSearchedPerson = make(map[string]bool)
}

// addPerson adds people to a stack or queue
func addPerson(pq *[]string, people []string) {
	for i := range people {
		*pq = append(*pq, people[i])
	}
}

func is_mango_seller(person string) bool {
	return person[len(person)-1] == 'm'
}

func bfs() bool {
	addPerson(&personQueen, netGraph["you"])
	fmt.Printf("you ")

	for len(personQueen) > 0 {
		p := personQueen[0]
		personQueen = personQueen[1:]
		if bfsSearchedPerson[p] {
			continue
		}
		fmt.Printf("-> %s", p)
		if is_mango_seller(p) {
			fmt.Printf("\n")
			fmt.Printf("%s is a mango seller\n", p)
			return true
		}
		bfsSearchedPerson[p] = true
		addPerson(&personQueen, netGraph[p])
	}
	return false
}

// add dfs implementation for practice
func dfs() bool {
	addPerson(&personStack, netGraph["you"])
	fmt.Printf("you ")

	for len(personStack) > 0 {
		p := personStack[len(personStack)-1]
		personStack = personStack[:len(personStack)-1]
		if dfsSearchedPerson[p] {
			continue
		}
		fmt.Printf("-> %s", p)
		if is_mango_seller(p) {
			fmt.Printf("\n")
			fmt.Printf("%s is a mango seller\n", p)
			return true
		}
		dfsSearchedPerson[p] = true
		addPerson(&personStack, netGraph[p])
	}
	return false
}

func main() {
	fmt.Println(bfs())
	fmt.Println(dfs())
}
