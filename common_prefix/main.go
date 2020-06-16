package main

import "fmt"

func main() {
	fmt.Println(
		longestCommonPrefix([]string{"flower", "flow", "flight"}),
	)

}

func longestCommonPrefix(strs []string) string {

	var prefix string
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	firstWord := strs[0]
END:
	for i := 0; i < len(firstWord); i++ {
		for _, str := range strs[1:] {
			if str[i] != firstWord[i] {
				break END
			}
		}
		prefix += string(firstWord[i])
	}

	return prefix
}
