// Example code for https://guo-sj.github.io/algorithm/2022/04/19/bfs.html
package main

import (
	"fmt"
	"log"
)

type netGraph map[string][]string

type personQueen []string

type searchedPerson map[string]bool

func (ng *netGraph) initGraph() {
	*ng = make(netGraph)
	(*ng)["you"] = []string{"alice", "bob", "claire"}
	(*ng)["bob"] = []string{"anuj", "peggy"}
	(*ng)["alice"] = []string{"peggy"}
	(*ng)["claire"] = []string{"thom", "jonny"}
	(*ng)["anuj"] = []string{""}
	(*ng)["peggy"] = []string{""}
	(*ng)["thom"] = []string{""}
	(*ng)["jonny"] = []string{""}
}

func (pq *personQueen) add(people []string) {
	log.Printf("add %v to pq\n", people)
	for i := range people {
		*pq = append(*pq, people[i])
	}
}

func (sp *searchedPerson) add(person string) {
	(*sp)[person] = true
}

func is_mango_seller(person string) bool {
	log.Printf("is_mango_seller: %s\n", person)
	return person[len(person)-1] == 'm'
}

func bfs() bool {
	var pq personQueen
	var ng netGraph
	ng.initGraph()
	pq.add(ng["you"])

	sq := make(searchedPerson)
	for i := 0; i < len(pq); i++ {
		p := pq[i]
		if sq[p] {
			continue
		}
		if is_mango_seller(p) {
			fmt.Printf("%s is a mango seller\n", p)
			return true
		}
		sq.add(p)
		pq.add(ng[p])
	}
	return false
}

func main() {
	fmt.Println(bfs())
}
