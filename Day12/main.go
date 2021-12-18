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
	input := read.ReadStrArrayByLine("./smallinput.txt")

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

func followNodesPt1(curNode Node, nodeMap map[string]Node, path []string) {
	//fmt.Printf("For path %#v\n", path)
	//fmt.Printf("FollowNode into %s\n", curNode.name)
	path = append(path, curNode.name)
	if curNode.name == "end" {
		fmt.Printf("Ending path %#v\n", path)
		pathsPt1 = append(pathsPt1, path)
	} else {
		for _, nodeName := range curNode.connections {
			//fmt.Printf("Checking if we can go to %s\n", nodeName)
			isPresentAlready := false
			for _, node := range path {
				if node == nodeName {
					isPresentAlready = true
				}
			}
			if !nodeMap[nodeName].allowRevisit && isPresentAlready {
				//do nothing
				//fmt.Printf("Terminating on node %s\n", curNode.name)
			} else {
				//fmt.Printf("Able to go to %s\n", nodeMap[nodeName].name)
				followNodesPt1(nodeMap[nodeName], nodeMap, path)
			}
		}
	}
}

func followNodesPt2(curNode Node, nodeMap map[string]Node, path []string, specialCave string, specialCaveVisits int) {
	//fmt.Printf("For path %#v\n", path)
	//fmt.Printf("FollowNode into %s\n", curNode.name)
	path = append(path, curNode.name)
	if curNode.name == "end" {
		fmt.Printf("Ending path %#v\n", path)
		pathsPt2 = append(pathsPt2, path)
		//fmt.Printf("CurPaths: %#v\n", pathsPt2)
	} else {
		for _, nodeName := range curNode.connections {
			//fmt.Printf("Checking if we can go to %s\n", nodeName)
			presentCount := 0
			for _, node := range path {
				if node == nodeName {
					presentCount = presentCount + 1
				}
			}
			if !nodeMap[nodeName].allowRevisit && (presentCount == 1 && nodeName != specialCave) {
				//do nothing
				fmt.Printf("    Failed path %#v >> %s\n", path, nodeName)
			} else if nodeMap[nodeName].allowRevisit || (nodeName == specialCave && presentCount <= 1) || nodeName == "end" {
				//fmt.Printf("Able to go to %s\n", nodeMap[nodeName].name)
				followNodesPt2(nodeMap[nodeName], nodeMap, path, specialCave, specialCaveVisits)
			}
		}
	}
}

var pathsPt1 [][]string
var pathsPt2 [][]string

func part1() {
	nodeMap := parseInput()
	//fmt.Printf("Created Nodes %#v\n", nodeMap)
	for _, nodeName := range nodeMap["start"].connections {
		path := []string{"start"}
		followNodesPt1(nodeMap[nodeName], nodeMap, path)
	}

	fmt.Printf("Part 1 - Number of paths: %d\n", len(pathsPt1))
}

func part2() {
	nodeMap := parseInput()
	//fmt.Printf("Created Nodes %#v\n", nodeMap)

	keys := make([]string, 0, len(nodeMap))
	for k := range nodeMap {
		keys = append(keys, k)
	}

	var smallCaves []string
	for _, name := range keys {
		small := strings.ContainsAny(name, "abcdefghijklmnopqrstuvwxyz")
		if small && name != "start" && name != "end" {
			smallCaves = append(smallCaves, name)
		}
	}

	for _, specialCave := range smallCaves {
		fmt.Printf("SpecialCave is %s\n", specialCave)
		for _, nodeName := range nodeMap["start"].connections {
			specialCaveVisits := 0
			path := []string{"start"}
			followNodesPt2(nodeMap[nodeName], nodeMap, path, specialCave, specialCaveVisits)
		}
	}

	var fixedPaths [][]string
	for _, path := range pathsPt2 {
		var curPath []string
		for _, node := range path {
			if node == "end" {
				curPath = append(curPath, node)
				break
			} else {
				curPath = append(curPath, node)
			}
		}
		//fmt.Printf("Fixed path %#v\n", curPath)
		fixedPaths = append(fixedPaths, curPath)
	}

	pathMap := make(map[string]int)
	for _, path := range fixedPaths {
		pathString := ""
		for _, node := range path {
			pathString = pathString + node
		}
		pathMap[pathString] = pathMap[pathString] + 1
	}

	fmt.Printf("Paths: %#v\n", pathMap)

	fmt.Printf("Part 2 - Number of paths: %d\n", len(pathMap))
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
