package main

import (
	"fmt"
	"strings"
)

/*
 * leetcode: https://leetcode-cn.com/problems/integer-to-roman/
 */

func main() {

	ret := intToRoman(1994)
	fmt.Println(ret)
}

func greedyIntToRoman(num int) string {

	return ""
}

func intToRoman(num int) string {
	if num == 0 {
		return ""
	}
	if num > 3999 || num < 1 {
		panic(`input out range of 1~3999`)
	}

	if num >= 1 && num < 10 {
		switch num {
		case 4:
			return "IV"
		case 9:
			return "IX"
		default:
			if num >= 5 {
				return "V" + strings.Repeat("I", num-5)
			} else {
				return strings.Repeat("I", num)
			}
		}
	} else if num >= 10 && num < 50 {
		if num >= 40 {
			return "XL" + intToRoman(num-40)
		} else {
			return strings.Repeat("X", num/10) + intToRoman(num%10)
		}
	} else if num >= 50 && num < 100 {
		if num >= 90 {
			return "XC" + intToRoman(num-90)
		} else {
			return "L" + intToRoman(num-50)
		}
	} else if num >= 100 && num < 500 {
		if num >= 400 {
			return "CD" + intToRoman(num-400)
		} else {
			return strings.Repeat("C", num/100) + intToRoman(num%100)
		}
	} else if num >= 500 && num < 1000 {
		if num >= 900 {
			return "CM" + intToRoman(num-900)
		} else {
			return "D" + intToRoman(num-500)
		}
	} else {
		return strings.Repeat("M", num/1000) + intToRoman(num%1000)
	}
}
