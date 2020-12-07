package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	bag  string
	reqs []required
}

type required struct {
	amount int
	bag    string
}

func main() {
	rules, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(findBag("shiny gold", rules))
}

func rulesToGraph(rules []rule) *Graph {
	g := NewGraph()
	for _, rule := range rules {
		g.AddNode(rule.bag)
	}
	for _, rule := range rules {
		for _, baginList := range rule.reqs {
			err := g.AddEdge(baginList.bag, rule.bag, baginList.amount)
			if err != nil {
				panic(err)
			}
		}
	}
	return g
}

func findBag(bag string, rules []rule) (count int) {
	g := rulesToGraph(rules)
	result, err := g.TopSort(bag)
	if err != nil {
		panic(err)
	}
	return len(result) - 1
}

func readInput(filename string) (rules []rule, err error) {
	rules = make([]rule, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return rules, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Split(line, " ")
		bag := token[0] + " " + token[1]
		r := rule{bag: bag}
		for i := 4; i < len(token); i += 4 {
			if token[i] == "no" {
				break
			}
			amount, err := strconv.Atoi(token[i])
			if err != nil {
				return rules, err
			}
			cur := token[i+1] + " " + token[i+2]
			r.reqs = append(r.reqs, required{amount, cur})
		}
		rules = append(rules, r)
	}
	return rules, nil
}

/* https://gostudent.github.io/Letsgo/Implementation-of-Topological-sort-in-GO */

type Graph struct {
	nodes map[string]node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]node),
	}
}

func (g *Graph) AddNode(name string) {
	if !g.ContainsNode(name) {
		g.nodes[name] = make(node)
	}
}

func (g *Graph) AddEdge(from string, to string, weight int) error {
	f, ok := g.nodes[from]
	if !ok {
		return fmt.Errorf("Node %q not found", from)
	}
	_, ok = g.nodes[to]
	if !ok {
		return fmt.Errorf("Node %q not found", to)
	}

	f.addEdge(to, weight)
	return nil
}

func (g *Graph) ContainsNode(name string) bool {
	_, ok := g.nodes[name]
	return ok
}

func (g *Graph) TopSort(name string) ([]string, error) {
	results := newOrderedSet()
	err := g.visit(name, results, nil)
	if err != nil {
		return nil, err
	}
	return results.items, nil
}

func (g *Graph) visit(name string, results *orderedset, visited *orderedset) error {
	if visited == nil {
		visited = newOrderedSet()
	}

	added := visited.add(name)
	if !added {
		index := visited.index(name)
		cycle := append(visited.items[index:], name)
		return fmt.Errorf("Cycle error: %s", strings.Join(cycle, " -> "))
	}

	n := g.nodes[name]
	for _, edge := range n.edges() {
		err := g.visit(edge, results, visited.copy())
		if err != nil {
			return err
		}
	}

	results.add(name)
	return nil
}

type node map[string]int

func (n node) addEdge(name string, weight int) {
	n[name] = weight
}

func (n node) edges() []string {
	var keys []string
	for k := range n {
		keys = append(keys, k)
	}
	return keys
}

type orderedset struct {
	indexes map[string]int
	items   []string
	length  int
}

func newOrderedSet() *orderedset {
	return &orderedset{
		indexes: make(map[string]int),
		length:  0,
	}
}

func (s *orderedset) add(item string) bool {
	_, ok := s.indexes[item]
	if !ok {
		s.indexes[item] = s.length
		s.items = append(s.items, item)
		s.length++
	}
	return !ok
}

func (s *orderedset) copy() *orderedset {
	clone := newOrderedSet()
	for _, item := range s.items {
		clone.add(item)
	}
	return clone
}

func (s *orderedset) index(item string) int {
	index, ok := s.indexes[item]
	if ok {
		return index
	}
	return -1
}
