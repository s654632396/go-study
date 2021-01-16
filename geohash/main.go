package main

import (
	"bytes"
	"fmt"
)

const (
	MaxLat float64 = 90
	MinLat float64 = -90
	MaxLng float64 = 180
	MinLng float64 = -180
	Length         = 20
)

var (
	base32Lookup = []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
		'j', 'k', 'm', 'n', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
)

func convert(min, max, value float64, list []byte) []byte {
	if len(list) >= Length {
		return list
	}
	var mid float64 = (max + min) / 2
	if value >= mid {
		list = append(list, '1')
		return convert(mid, max, value, list)
	} else {
		list = append(list, '0')
		return convert(min, mid, value, list)
	}

}

func combineLatLng(lat, lng []byte) []byte {
	var bf bytes.Buffer
	for i := 0; i < len(lat); i++ {
		bf.WriteByte(lng[i])
		bf.WriteByte(lat[i])
	}
	return bf.Bytes()
}

func encodeToIndex(list []byte) string {
	var bf bytes.Buffer

	for i := 0; i < len(list); i += 5 {
		var unit int = 0
		for j := i; j < i+5; j++ {
			if list[j] == '1' {
				unit += 1 << (4 - (j - i))
			}
		}
		bf.WriteByte(base32Lookup[unit])
	}

	return bf.String()
}

func main() {
	lat, lng := 39.92324, 116.3906
	latList := make([]byte, 0)
	lngList := make([]byte, 0)
	latList = convert(MinLat, MaxLat, lat, latList)
	lngList = convert(MinLng, MaxLng, lng, lngList)
	fmt.Println(string(latList), string(lngList))
	list := combineLatLng(latList, lngList)
	fmt.Println(string(list))
	index := encodeToIndex(list)
	fmt.Println(index)
}
