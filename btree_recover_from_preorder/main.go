package main

import (
	"bytes"
	"fmt"
	"strconv"
)

/**
 * https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/
 */
func main() {

	fmt.Println(
		serialize(
			//recoverFromPreorder("1-2--3--4-5--6--7"),
			//recoverFromPreorder("1-2--3---4-5--6---7"),
			//recoverFromPreorder("1-401--349---90--88"),
			recoverFromPreorder("8-6--9---10--4-7"),
		),
	)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(S string) *TreeNode {
	var (
		tree           *TreeNode              // 需要返回的Tree Root Node
		currentNode    string                 // 节点字符串
		level          int                    // 深度
		lastLevel      = level                // 上一次的深度
		nodeStack      = make([]*TreeNode, 1) // 节点堆
		nodeStackLevel = make([]int, 1)       // 节点堆深度
	)

	for i := 0; i <= len(S); i++ {
		if (i == len(S) || S[i] == '-') && len(currentNode) > 0 {
			// 初始化即将插入树上的节点
			tmpInt, _ := strconv.Atoi(currentNode)
			node := &TreeNode{tmpInt, nil, nil}
			var parentNode *TreeNode

			if nodeStack[0] == nil {
				// Tree的根
				nodeStack[0] = node
				tree = node
				nodeStackLevel[0] = 0
			} else {
				// 找父节点
				if lastLevel < level {
					// pre-order向下遍历, 父节点一定是在堆的上一个
					parentNode = nodeStack[len(nodeStack)-1]
				} else {
					// 从堆上查找上一层深度的父节点
					parentNode = nil
					for j := len(nodeStack) - 1 - (lastLevel - level + 1); j >= 0; j-- {
						if nodeStackLevel[j] == level-1 {
							parentNode = nodeStack[j]
							break
						}
					}
				}
				if parentNode == nil {
					panic(`could not find parentNode`)
				}
				// 把节点插到父节点下
				if parentNode.Left == nil {
					parentNode.Left = node
				} else {
					parentNode.Right = node
				}
				// 把节点追加到节点堆上,并记录深度
				nodeStack = append(nodeStack, node)
				nodeStackLevel = append(nodeStackLevel, level)
			}

			// reset
			currentNode = ""  // 清空
			lastLevel = level // 记录上一个append node的level
			level = 1         // 深度回归初始(=1的原因是当前已经匹配到一个"-",最后一个元素除外)
		} else if S[i] == '-' {
			// 每匹配到一个"-", 当前元素的深度+1
			level += 1
			continue
		} else {
			currentNode += string(S[i])
		}
	}
	return tree
}

// Serializes a tree to a single string.
func serialize(root *TreeNode) string {

	var nodeMap = make(map[int][]*TreeNode)
	nodeMap[1] = []*TreeNode{root}

	var ret = make([]string, 0)

	for l := 1; ; l++ {
		cNodeList := nodeMap[l]
		var hasNext = false
		if nodeMap[l+1] == nil {
			nodeMap[l+1] = make([]*TreeNode, 0)
		}
		for _, node := range cNodeList {
			var val string
			if node != nil {
				val = strconv.Itoa(node.Val)
			} else {
				val = "null"
			}
			ret = append(ret, val)
			if node == nil {
				continue
			}
			if !hasNext && (node.Left != nil || node.Right != nil) {
				hasNext = true
			}
			nodeMap[l+1] = append(nodeMap[l+1], node.Left)
			nodeMap[l+1] = append(nodeMap[l+1], node.Right)
		}
		if !hasNext {
			break
		}
	}
	var bf bytes.Buffer
	bf.WriteString("[")
	for i := 0; i < len(ret); i++ {
		bf.WriteString(ret[i])
		if i != len(ret)-1 {
			bf.WriteString(",")
		}
	}
	bf.WriteString("]")
	return bf.String()
}
