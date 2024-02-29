/*
Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

```
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
```

The following test should pass:

```
node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'
```

*/

package main

import (
	"fmt"
	"regexp"
)

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func main() {
	node := Node{Val: "roo\\\"t", Left: &Node{Val: "left", Left: &Node{Val: "left.left"}}, Right: &Node{Val: "right"}}
	serializedNode := serialize(&node)
	fmt.Println(deserialize(serializedNode).Left.Left.Val)
}

func serialize(node *Node) string {
	serializedLeft := ""
	serializedRight := ""
	if node.Left != nil {
		serializedLeft = serialize(node.Left)
	}
	if node.Right != nil {
		serializedRight = serialize(node.Right)
	}
	return "(\"" + escape(node.Val) + "\"," + serializedLeft + "," + serializedRight + ")"
}

func escape(value string) string {
	quotePattern := regexp.MustCompile("\"")
	backslashPattern := regexp.MustCompilePOSIX("\\\\")
	modified := backslashPattern.ReplaceAll([]byte(value), []byte("\\\\"))
	modified = quotePattern.ReplaceAll(modified, []byte("\\\""))
	return string(modified)
}

func unescape(value string) string {
	quotePattern := regexp.MustCompile("\\\\\"")
	backslashPattern := regexp.MustCompilePOSIX("\\\\\\\\")
	modified := quotePattern.ReplaceAll([]byte(value), []byte("\""))
	modified = backslashPattern.ReplaceAll(modified, []byte("\\"))
	return string(modified)
}

func deserialize(serializedNode string) *Node {
	if serializedNode == "" {
		return nil
	} else {
		parameters := parseParameters(serializedNode)
		return &Node{Val: unescape(parameters[0]), Left: deserialize(parameters[1]), Right: deserialize(parameters[2])}
	}
}

func parseParameters(serializedNode string) [3]string {
	parameters := [3]string{}
	index := 0
	splitPosition := 0
	serializedNode = serializedNode[1 : len(serializedNode)-1]
	var inQuotes = false
	var parenthesesCount = 0
	var escapeNextCharacter = false
	for i := 0; i < len(serializedNode); i++ {
		if escapeNextCharacter {
			escapeNextCharacter = false
		} else if serializedNode[i] == '"' && parenthesesCount == 0 {
			inQuotes = !inQuotes
		} else if serializedNode[i] == ',' && !inQuotes && parenthesesCount == 0 {
			parameters[index] = serializedNode[splitPosition:i]
			index += 1
			splitPosition = i + 1
		} else if serializedNode[i] == '(' && !inQuotes {
			parenthesesCount += 1
		} else if serializedNode[i] == ')' && !inQuotes {
			parenthesesCount -= 1
		} else if serializedNode[i] == '\\' && inQuotes {
			escapeNextCharacter = true
		}
		if i == len(serializedNode)-1 {
			parameters[index] = serializedNode[splitPosition:]
		}
	}
	return parameters
}
