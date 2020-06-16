package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(
		simplifyPath("/a/../../b/../c//.//"),
		simplifyPath("/a//b////c/d//././/.."),
	)

}

func simplifyPath(path string) string {
	var pathArr = make([]string, 0)

	ori := strings.Split(path, "/")

	for _, dir := range ori {
		if dir == "" || dir == "." {
			continue
		}
		if dir == ".." {
			if len(pathArr)-1 < 0 {
				pathArr = []string{}
			} else {
				pathArr = pathArr[:len(pathArr)-1]
			}
			continue
		}
		pathArr = append(pathArr, dir)
	}

	return "/" + strings.Join(pathArr, "/")
}
