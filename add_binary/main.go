package main

import (
	"bytes"
	"fmt"
)

func main() {
	//r:=addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011")
	r := addBinary("11", "1")
	fmt.Println(r)
}

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)

	var l int
	if la > lb {
		l = la
	} else {
		l = lb
	}
	arrA := make([]bool, l)
	arrB := make([]bool, l)

	var arrC = make([]bool, l+1)
	var plusF bool
	var bf bytes.Buffer

	for i := 0; i < l; i++ {
		if plusF {
			arrC[i] = true
			plusF = false
		}

		if i+1 <= la && a[la-1-i] == '1' {
			arrA[i] = true
		}
		if i+1 <= lb && b[lb-1-i] == '1' {
			arrB[i] = true
		}
		// -----calculate-----
		if arrA[i] && arrB[i] {
			plusF = true
		} else if (arrA[i] || arrB[i]) && arrC[i] {
			plusF = true
			arrC[i] = false
		} else if !arrC[i] && (arrA[i] || arrB[i]) {
			arrC[i] = true
		}
		sub := bf.String()
		if arrC[i] {
			bf.Reset()
			bf.WriteString("1" + sub)
		} else {
			bf.Reset()
			bf.WriteString("0" + sub)
		}
	}
	if plusF {
		arrC[len(arrC)-1] = true
		sub := bf.String()
		bf.Reset()
		bf.WriteString("1" + sub)
	}

	return bf.String()
}
