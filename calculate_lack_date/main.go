package main

import (
	"log"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type NodeFlag int8

const (
	NodeFlagStart NodeFlag = 0
	NodeFlagEnd   NodeFlag = 1
)

type dateNode struct {
	Year  int
	Month int // 1~12
	F     NodeFlag
}

type dateNodes []dateNode

func (ns dateNodes) Len() int {
	return len(ns)
}

func (ns dateNodes) Less(i, j int) bool {
	if ns[i].Year < ns[j].Year {
		return true
	} else
	if ns[i].Year > ns[j].Year {
		return false
	} else {
		// equal of year, compare month
		if ns[i].Month <= ns[j].Month {
			return true
		} else {
			return false
		}
	}
}

func (ns dateNodes) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func main() {
	// must be reverse sorted by start_time first
	var data = []map[string]string{
		{"start_time": "2018年09月", "end_time": "", "so_far": "Y"},
		{"start_time": "2014年09月", "end_time": "2018年08月", "so_far": "N"},
		{"start_time": "2010年07月", "end_time": "2014年08月", "so_far": "N"},
		{"start_time": "2006年05月", "end_time": "2010年04月", "so_far": "N"},
	}

	nodes := parseToNode(data)
	log.Println(nodes)

	var flag = false
	var cStart dateNode
	var monthes int
	for i := 0; i < nodes.Len(); i++ {
		if nodes[i].F == NodeFlagEnd {
			flag = true
			cStart = nodes[i]
		} else {
			if flag {
				//do calculate
				monthes += calculateDiffMonthBetweenNode(cStart, nodes[i])
				flag = false
			}
		}
	}

	log.Println("empty monthes=", monthes)
}

//
func calculateDiffMonthBetweenNode(start, end dateNode) int {
	var m = 0
	var ceilYear = end.Year - start.Year
	m = ceilYear*12 + end.Month - start.Month
	return m
}

func parseToNode(data []map[string]string) dateNodes {
	var nodes = make(dateNodes, 0)
	reg, err := regexp.Compile(`(\d{4})年(\d{1,2})月`)
	if err != nil {
		log.Println(err)
	}
	for i:=len(data)-1;i >=0; i-- {
		var s = data[i]
		//get start node
		var node = new(dateNode)
		node.F = NodeFlagStart
		var matches = reg.FindAllStringSubmatch(s["start_time"], 1)
		node.Year, _ = strconv.Atoi(matches[0][1])
		node.Month, _ = strconv.Atoi(matches[0][2])
		nodes = append(nodes, *node)
		if nodes.Len() > 1 && false == nodes.Less(nodes.Len()-2, nodes.Len()-1) {
			nodes[nodes.Len()-2].F = NodeFlagStart
		}
		//get end node
		node = new(dateNode)
		node.F = NodeFlagEnd
		if s["so_far"] != "N" {
			node.Year = time.Now().Year()
			node.Month = int(time.Now().Month())
		} else {
			matches = reg.FindAllStringSubmatch(s["end_time"], 1)
			node.Year, _ = strconv.Atoi(matches[0][1])
			node.Month, _ = strconv.Atoi(matches[0][2])
		}
		nodes = append(nodes, *node)
	}
	sort.Sort(nodes)

	return nodes
}
