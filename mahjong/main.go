package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// 写着玩玩
// 日麻牌理
// 定义 万 = m | 饼 = p | 索 = s | 东南西北白发中 = 1234567z
// 定义 红5 = 0
var CardMap = map[string]int64{
	"1z": 0x1F000,
	"2z": 0x1F001,
	"3z": 0x1F002,
	"4z": 0x1F003,
	"5z": 0x1F004,
	"6z": 0x1F005,
	"7z": 0x1F006,
	"1m": 0x1F007,
	"2m": 0x1F008,
	"3m": 0x1F009,
	"4m": 0x1F00A,
	"5m": 0x1F00B,
	"6m": 0x1F00C,
	"7m": 0x1F00D,
	"8m": 0x1F00E,
	"9m": 0x1F00F,
	"1s": 0x1F010,
	"2s": 0x1F011,
	"3s": 0x1F012,
	"4s": 0x1F013,
	"5s": 0x1F014,
	"6s": 0x1F015,
	"7s": 0x1F016,
	"8s": 0x1F017,
	"9s": 0x1F018,
	"1p": 0x1F019,
	"2p": 0x1F01A,
	"3p": 0x1F01B,
	"4p": 0x1F01C,
	"5p": 0x1F01D,
	"6p": 0x1F01E,
	"7p": 0x1F01F,
	"8p": 0x1F020,
	"9p": 0x1F021,
}

type 张类型 string

const (
	万 张类型 = "m"
	索 张类型 = "s"
	饼 张类型 = "p"
	字 张类型 = "z"
)

type 张 struct {
	num int // range of 1~9
	t   张类型
}

func (z *张) String() string {
	n := strconv.Itoa(z.num)
	return n + string(z.t)
}

var (
	听牌枚数 = []int{4, 7, 10, 13}
	和牌枚数 = []int{5, 8, 11, 14}
)

type 听牌类型 string

const (
	两面 听牌类型 = "1"
	坎张 听牌类型 = "2"
	边听 听牌类型 = "3"
	钓将 听牌类型 = "4"
)

func main() {
	var handCardstring string = "123m123p22334s11z"
	var handCard = ReadHandCard(handCardstring)
	// fmt.Printf("handCard is : %+v\n", handCard)
	ToUnicode(handCard)
	ReadTallyCount(handCard)

}

// ToUnicode Output of Unicode
func ToUnicode(hc []张) {
	for _, item := range hc {
		fmt.Print(string(CardMap[item.String()]), " ")
	}
	fmt.Println()
}

// ReadHandCard  读取手牌字符串，转换为手牌数组
func ReadHandCard(hcs string) (ret []张) {

	re, _ := regexp.Compile("[0-9]{1,14}[m|p|s|z]")
	// fmt.Println(re.MatchString(hcs))
	var cardSlice []string = re.FindAllString(hcs, -1)
	for _, v := range cardSlice {
		var pieceLen int = len(v)
		pieceNum, pieceMark := v[:pieceLen-1], v[pieceLen-1:]
		for _, n := range []rune(pieceNum) {
			number, _ := strconv.Atoi(string(n))
			zhang := 张{number, (张类型)(pieceMark)}
			ret = append(ret, zhang)
		}
	}

	return
}

// ReadTallyCount 读取向听数 max = 6, min = 0
func ReadTallyCount(hc []张) (tc int) {
	// thirteenTallyCount := getThirteenTallyCount(hc)
	// fmt.Printf("国士无双向听听数为: %d \n", thirteenTallyCount)

	// sevenPairsCount := getSevenPairsCount(hc)
	// fmt.Printf("七对子向听数为: %d \n", sevenPairsCount)

	nornalTallyCount := getNormalTallyCount(hc)
	fmt.Printf("普通牌型向听数为: %d \n", nornalTallyCount)
	return
}

// getNormalTallyCount 取正常手顺牌型向听
func getNormalTallyCount(hc []张) (ret int) {
	ret = 4
	// fmt.Println(hc)
	// 和牌牌型要满足 3n * 4 + 2p的公式
	// 3n 是连续的3张数牌或者相同的牌
	// 2p 是任意相同的2枚牌

	// 如果定义听牌向听数 X = 0
	// 有且仅有一个2p时，算  - 1
	// 每个序列每满足一个3n时， 算 - 1
	// 则 Ymax - 1 * 4 = X ,  Ymax = 4

	// fmt.Println(hc)
	// var itc = 4 // 起始向听数 ***再大的向听也记为4向听
	//
	var (
		typeM []int
		typeP []int
		typeS []int
		typeZ []int
	)
	for _, c := range hc {
		cardNum, cardType := c.num, c.t

		switch cardType {
		case 万:
			typeM = append(typeM, cardNum)
		case 饼:
			typeP = append(typeP, cardNum)
		case 索:
			typeS = append(typeS, cardNum)
		case 字:
			typeZ = append(typeZ, cardNum)
		}
	}
	fmt.Printf("typeM=%+v, typeP=%+v, typeS=%+v, typeZ=%+v \n", typeM, typeP, typeS, typeZ)
	return
}

func calculate(cards []int) {

}

// ------------
type structMap map[string]int

// UniqueTypeCardCollate 唯一牌组集合
type UniqueTypeCardCollate struct {
	cardMap structMap
}

func (utcc *UniqueTypeCardCollate) pushCard(card string) {
	if utcc.cardMap == nil {
		utcc.cardMap = make(structMap)
	}
	utcc.cardMap[card]++
}

func (utcc *UniqueTypeCardCollate) existsCard(card string) bool {
	var ret = false
	if _, ok := utcc.cardMap[card]; ok {
		ret = true
	}
	return ret
}

// 数对子
func (utcc *UniqueTypeCardCollate) coupleCount() int {
	var count int = 0
	for _, v := range utcc.cardMap {
		if v >= 2 {
			count++
		}
	}

	return count
}

// getThirteenTallyCount 取国士牌型向听数 ret ~ [0, 13]
func getThirteenTallyCount(hc []string) (ret int) {
	// 国士无双的牌型，先判断共有多少个单张的19字牌，然后判断是否有至少一个19字牌的对子
	ret = 13
	judgeCard := []string{"1m", "9m", "1p", "9p", "1s", "9s", "1z", "2z", "3z", "4z", "5z", "6z", "7z"}

	var hasCoupleCard = false
	// 简单粗暴双循环
	var cardSet UniqueTypeCardCollate
	for _, card := range judgeCard {
		for _, handCard := range hc {
			if handCard == card {
				if cardSet.existsCard(card) == false {
					cardSet.pushCard(card)
					ret--
				} else {
					hasCoupleCard = true
				}
			}
		}
	}
	if hasCoupleCard {
		ret--
	}
	return
}

// getSevenPairsCount 取七对子牌型向听数 ret ~[0, 6]
func getSevenPairsCount(hc []string) (ret int) {
	// 七对子的向听，通过多少对来判断，切不能有重复对子

	var cardSet UniqueTypeCardCollate
	for _, card := range hc {
		cardSet.pushCard(card)
	}

	ret = 6 - cardSet.coupleCount()

	return
}