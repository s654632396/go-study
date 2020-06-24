package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
)

// 写着玩玩
// 日麻牌理
// 定义 TypeM = m | TypeP = p | TypeS = s | 东南西北白发中 = 1234567z
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

type CardType string

const (
	TypeM CardType = "m"
	TypeP CardType = "p"
	TypeS CardType = "s"
	TypeZ CardType = "z"
)
const (
	PadTypeM int = 0
	PadTypeP int = 1
	PadTypeS int = 2
	PadTypeZ int = 3
)

type Card struct {
	num int // range of 1~9
	t   CardType
}

func (z *Card) String() string {
	n := strconv.Itoa(z.num)
	return n + string(z.t)
}

var (
	WaitNumber = []int{4, 7, 10, 13}
	PushNumber = []int{5, 8, 11, 14}
)

type WaitType string

const (
	WaitTwoPiece  WaitType = "1" // 两面
	WaitInPiece   WaitType = "2" // 坎张
	WaitSidePiece WaitType = "3" // 边张
	WaitPairPiece WaitType = "4" // 单骑
)

func main() {
	var handCardstring string = "889123m23234s67z"
	var handCard = ReadHandCard(handCardstring)
	// fmt.Printf("handCard is : %+v\n", handCard)
	ToUnicode(handCard)
	ReadTallyCount(handCard)

}

// ToUnicode Output of Unicode
func ToUnicode(hc []Card) {
	for _, item := range hc {
		fmt.Print(string(CardMap[item.String()]), " ")
	}
	fmt.Println()
}

// ReadHandCard  读取手牌字符串，转换为手牌数组
func ReadHandCard(hcs string) (ret []Card) {

	re, _ := regexp.Compile("[0-9]{1,14}[m|p|s|z]")
	// fmt.Println(re.MatchString(hcs))
	var cardSlice []string = re.FindAllString(hcs, -1)
	for _, v := range cardSlice {
		var pieceLen int = len(v)
		pieceNum, pieceMark := v[:pieceLen-1], v[pieceLen-1:]
		for _, n := range []rune(pieceNum) {
			number, _ := strconv.Atoi(string(n))
			zhang := Card{number, (CardType)(pieceMark)}
			ret = append(ret, zhang)
		}
	}

	return
}

// ReadTallyCount 读取向听数 max = 6, min = 0
func ReadTallyCount(hc []Card) (tc int) {
	// thirteenTallyCount := getThirteenTallyCount(hc)
	// fmt.Printf("国士无双向听听数为: %d \n", thirteenTallyCount)

	// sevenPairsCount := getSevenPairsCount(hc)
	// fmt.Printf("七对子向听数为: %d \n", sevenPairsCount)

	_ = getNormalTallyCount(hc)
	// fmt.Printf("普通牌型向听数为: %d \n", nornalTallyCount)
	return
}

// getNormalTallyCount 取正常手顺牌型向听
func getNormalTallyCount(hc []Card) (ret int) {
	ret = 4

	calculate(hc)

	return
}

func calculate(cards []Card) {
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
	cardMatrix := getGroup(cards)
	// log.Println(cardMatrix)
	// 尽可能组合所有可能的拆解
	//_, shared := pickUnsharedCard([]int{1,1,1,2,3,4,5,5,5})
	hands := make([]int, len(cardMatrix[0])+len(cardMatrix[1])+len(cardMatrix[2])+len(cardMatrix[3]))
	// 1~9 		--- M
	// 11~19 	--- P
	// 21~29	--- S
	// 31~37	--- Z

	for i, padFix := 0, 0; i < len(hands); {
		for _, card := range cardMatrix[padFix] {
			hands[i] = padFix*10 + card
			i++
		}
		padFix++
	}
	log.Println(hands)
	soloCards, sharedCards := pickUnsharedCard(hands)
	log.Println(soloCards, sharedCards)
	inlineSharedCards(sharedCards)
}

type CardShunIface interface{}
type CardShun []int
type KeziShun CardShun
type SeqShun CardShun
type DaziShun CardShun
type QuetouShun CardShun

func inlineSharedCards(sharedCards []int) {
	log.Println("sharedCards=", sharedCards)
	// combines := make([][]CardShunIface, 0)
	hcMap := make(map[int]int)
	for _, x := range sharedCards {
		hcMap[x] ++
	}
	var cpHcMap = make(map[int]int)
	var uniques = make([]int, 0)
	for k, v := range hcMap {
		cpHcMap[k] = v
		uniques = append(uniques, k)
	}

	var counter = len(sharedCards)

	var combine = make([]CardShunIface, 0)
	for ; counter >= 0; counter-- {
		var hasQuetou bool
		// calculate
	NewCardShun:
		for _, x := range uniques {
			if cpHcMap[x] <= 0 {
				continue
			}
			cpHcMap[x]--
			padFix := x / 10
			lx, rx := max(x-2, padFix*10+1), min(x+2, padFix*10+9)
			for _, y := range uniques {
				if cpHcMap[y] <= 0 {
					continue
				}
				if y < lx || y > rx {
					continue
				}
				cpHcMap[y]--
				exps := parseDazi(x, y, padFix)
				for _, z := range exps {
					if _, ok := cpHcMap[z]; !ok {
						continue
					}
					if cpHcMap[z] == 0 {
						continue
					}
					cpHcMap[z]--
					if x == y {
						combine = append(combine, KeziShun{x, y, z})
					} else {
						combine = append(combine, SeqShun{x, y, z})
					}
					log.Println("matched:", x, y, z, cpHcMap)
					continue NewCardShun
				}
				// no matched z
				if x == y && !hasQuetou {
					combine = append(combine, QuetouShun{x, y})
				}
				if x == y && hasQuetou {
					combine = append(combine, DaziShun{x, y})
				}
				if x != y {
					combine = append(combine, DaziShun{x, y})
				}
				continue NewCardShun
			}
			// no matched ---
			cpHcMap[x]++
		}
	}
	log.Println(combine, cpHcMap)
}

func parseDazi(x, y int, padFix int) (expects []int) {
	if x == y {
		expects = []int{x}
		return
	}
	if x > y {
		x, y = y%10, x%10
	}
	if x+1 == y {
		if x == 1 {
			expects = []int{padFix*10 + 3}
		} else if x == 8 {
			expects = []int{padFix*10 + 7}
		} else {
			expects = []int{padFix*10 + x - 1, padFix*10 + y + 1}
		}
	} else if x+2 == y {
		if x == 1 {
			expects = []int{padFix*10 + 2}
		} else if x == 7 {
			expects = []int{padFix*10 + 8}
		} else {
			expects = []int{padFix*10 + x + 1}
		}
	} else {
		panic(`unknown err occurred`)
	}
	return
}

func getGroup(cards []Card) (cardMatrix [][]int) {
	cardMatrix = make([][]int, 4) // 4列, 依次为M,P,S,Z
	cardMatrix[0] = make([]int, 0)
	cardMatrix[1] = make([]int, 0)
	cardMatrix[2] = make([]int, 0)
	cardMatrix[3] = make([]int, 0)
	for _, c := range cards {
		cardNum, cardType := c.num, c.t
		switch cardType {
		case TypeM:
			cardMatrix[0] = append(cardMatrix[0], cardNum)
		case TypeP:
			cardMatrix[1] = append(cardMatrix[1], cardNum)
		case TypeS:
			cardMatrix[2] = append(cardMatrix[2], cardNum)
		case TypeZ:
			cardMatrix[3] = append(cardMatrix[3], cardNum)
		}
	}
	// 简单点,先排个序
	sort.Ints(cardMatrix[0])
	sort.Ints(cardMatrix[1])
	sort.Ints(cardMatrix[2])
	sort.Ints(cardMatrix[3])
	return
}

// 挑出孤张
func pickUnsharedCard(list []int) (unshared, shared []int) {
	if len(list) == 1 {
		unshared = list
	}
	for i := 0; i < len(list); i++ {
		var flag bool
		padFix := list[i] / 10
		var l, r int     // 计算靠张边界
		if padFix == 3 { // 字牌靠张等于自己
			r, l = list[i], list[i]
		} else {
			l, r = max(list[i]-2, padFix*10+1), min(list[i]+2, padFix*10+9)
		}
		for j := 0; j < len(list); j++ {
			if i == j {
				continue
			}
			if list[j] >= l && list[j] <= r {
				shared = append(shared, list[i])
				flag = true
				break
			}
		}
		if !flag {
			unshared = append(unshared, list[i])
		}
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// 是否对子
func isPairs(x, y int) bool {
	return x == y
}

// 是否刻子
func isColumnPair(x, y, z int) bool {
	return x == y && y == z
}

// 是否顺子
// x,y,z must be sorted
func isSequencePair(x, y, z int) bool {
	return x+1 == y && y+1 == z
}

// ------------
type structMap map[string]int

// UniqueTypeCardCollate 唯一牌组集合
type UniqueTypeCardCollate struct {
	cardMap structMap
}

func (cc *UniqueTypeCardCollate) SetCard(card string) {
	if cc.cardMap == nil {
		cc.cardMap = make(structMap)
	}
	cc.cardMap[card]++
}

func (cc *UniqueTypeCardCollate) existsCard(card string) bool {
	var ret = false
	if _, ok := cc.cardMap[card]; ok {
		ret = true
	}
	return ret
}

// 数对子
func (cc *UniqueTypeCardCollate) coupleCount() int {
	var count int = 0
	for _, v := range cc.cardMap {
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
					cardSet.SetCard(card)
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
		cardSet.SetCard(card)
	}

	ret = 6 - cardSet.coupleCount()

	return
}
