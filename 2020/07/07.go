package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Edge -
type Edge struct {
	nodeKey string
	weight  int
}

// Node -
type Node struct {
	key     string
	parents []Edge
}

func (node *Node) addParent(edge Edge) {
	node.parents = append(node.parents, edge)
}

func getKey(shade string, color string) string {
	return shade + " " + color
}

func getNode(key string, nodes map[string]*Node) *Node {
	if _, ok := nodes[key]; !ok {
		nodes[key] = &Node{key: key, parents: []Edge{}}
	}
	return nodes[key]
}

func parseRules(ruleDefinitions []string) map[string]*Node {
	nodes := map[string]*Node{}

	for _, rule := range ruleDefinitions {
		split := strings.Split(rule, " bags contain ")
		parentBagText := split[0]
		childBagsText := strings.Split(split[1], ", ")

		parentBagSplitText := strings.Split(parentBagText, " ")
		parentBagKey := getKey(parentBagSplitText[0], parentBagSplitText[1])
		parentBagNode := getNode(parentBagKey, nodes)

		if childBagsText[0] != "no other bags." {
			for _, child := range childBagsText {
				childBagSplitText := strings.Split(child, " ")
				amount, _ := strconv.Atoi(childBagSplitText[0])
				childKey := getKey(childBagSplitText[1], childBagSplitText[2])
				childNode := getNode(childKey, nodes)

				childNode.addParent(Edge{weight: amount, nodeKey: parentBagNode.key})
			}
		}
	}

	return nodes
}

func findParentBags(graph map[string]*Node, bagKey string) int {
	root := graph[bagKey]
	visited := map[string]bool{}
	traverse(graph, root, visited, 0)
	return len(visited) - 1
}

func traverse(graph map[string]*Node, node *Node, visited map[string]bool, depth int) {
	visited[node.key] = true

	for i := 0; i < len(node.parents); i++ {
		parent := node.parents[i]
		traverse(graph, graph[parent.nodeKey], visited, depth+1)
	}
}

func main() {
	lines := readLines("07.txt")
	rules := parseRules(lines)
	targetBagKey := "shiny gold"
	println(findParentBags(rules, targetBagKey))
}
