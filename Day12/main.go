package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

type Node struct {
	name         string
	connections  []string
	allowRevisit bool
}

func parseInput() map[string]Node {
	input := read.ReadStrArrayByLine("./exinput.txt")

	var nodes []Node

	for _, line := range input {
		split := strings.Split(line, "-")
		for curNode := 0; curNode < 2; curNode++ {
			var connectionNode string
			if curNode == 0 {
				connectionNode = split[1]
			} else {
				connectionNode = split[0]
			}
			foundExisting := false
			for nodeSearch := 0; nodeSearch < len(nodes); nodeSearch++ {
				if nodes[nodeSearch].name == split[curNode] {
					nodes[nodeSearch].connections = append(nodes[nodeSearch].connections, connectionNode)
					foundExisting = true
				}
			}
			if !foundExisting {
				allow := strings.ContainsAny(split[curNode], "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
				newNode := Node{
					name:         split[curNode],
					connections:  []string{connectionNode},
					allowRevisit: allow,
				}
				nodes = append(nodes, newNode)
			}
		}
	}

	nodeMap := make(map[string]Node)
	for _, node := range nodes {
		nodeMap[node.name] = node
	}

	return nodeMap
}

func followNodes(curNode Node, nodeMap map[string]Node) [][]string {
	for _, nodeName := range curNode.connections {

	}
}

func part1() {
	input := parseInput()
	//fmt.Printf("Created Nodes %#v\n", input)
	validPaths := followNodes(input["start"], input)
}

func part2() {

}

func main() {
	p1b := time.Now()
	part1()
	mid := time.Now()
	part2()
	p2a := time.Now()
	part1Time := mid.Sub(p1b)
	part2Time := p2a.Sub(mid)
	fmt.Printf("Part 1 Time: %vμs\nPart 2 Time: %vμs\n", part1Time.Microseconds(), part2Time.Microseconds())
}
