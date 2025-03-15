package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 游戏图标定义
const (
	GAME_SOLT_EMPTY = iota
	GAME_SOLT_1
	GAME_SOLT_2
	GAME_SOLT_3
	GAME_SOLT_4
	GAME_SOLT_5
	GAME_SOLT_6
	GAME_SOLT_7
	GAME_SOLT_8
	GAME_SOLT_WILD    //百搭
	GAME_SOLT_BOMB    //炸弹
	GAME_SOLT_ROCKET  //火箭
	GAME_SOLT_MINE    //地雷
	GAME_SOLT_SCATTER //scatter
	GAME_SOLT_MAX
)

const (
	//6行
	ROW_DEF = 6
	//6列
	COL_DEF = 6
	//图案元素
	BUFF_SIZE = ROW_DEF * COL_DEF
)

// 赔率表
var Mult = map[int]map[int]int{
	GAME_SOLT_1: {
		5:  2,
		6:  3,
		7:  4,
		8:  5,
		9:  6,
		10: 8,
		11: 10,
		12: 12,
		13: 14,
		14: 16,
		15: 20,
		16: 35,
		17: 40,
		18: 45,
		19: 50,
		20: 55,
		21: 60,
		22: 65,
		23: 70,
		24: 80,
		25: 150,
		26: 160,
		27: 170,
		28: 180,
		29: 190,
		30: 200,
		31: 220,
		32: 240,
		33: 260,
		34: 280,
		35: 300,
		36: 300,
	},
	GAME_SOLT_2: {
		5:  2,
		6:  3,
		7:  4,
		8:  5,
		9:  6,
		10: 8,
		11: 10,
		12: 12,
		13: 14,
		14: 16,
		15: 20,
		16: 35,
		17: 40,
		18: 45,
		19: 50,
		20: 55,
		21: 60,
		22: 65,
		23: 70,
		24: 80,
		25: 150,
		26: 160,
		27: 170,
		28: 180,
		29: 190,
		30: 200,
		31: 220,
		32: 240,
		33: 260,
		34: 280,
		35: 300,
		36: 300,
	},
	GAME_SOLT_3: {
		5:  3,
		6:  4,
		7:  5,
		8:  6,
		9:  8,
		10: 10,
		11: 12,
		12: 14,
		13: 16,
		14: 18,
		15: 25,
		16: 45,
		17: 60,
		18: 65,
		19: 70,
		20: 75,
		21: 80,
		22: 85,
		23: 90,
		24: 95,
		25: 180,
		26: 190,
		27: 200,
		28: 220,
		29: 240,
		30: 260,
		31: 280,
		32: 300,
		33: 320,
		34: 340,
		35: 360,
		36: 360,
	},
	GAME_SOLT_4: {
		5:  3,
		6:  4,
		7:  5,
		8:  6,
		9:  8,
		10: 10,
		11: 12,
		12: 14,
		13: 16,
		14: 18,
		15: 25,
		16: 45,
		17: 60,
		18: 65,
		19: 70,
		20: 75,
		21: 80,
		22: 85,
		23: 90,
		24: 95,
		25: 180,
		26: 190,
		27: 200,
		28: 220,
		29: 240,
		30: 260,
		31: 280,
		32: 300,
		33: 320,
		34: 340,
		35: 360,
		36: 360,
	},
	GAME_SOLT_5: {
		5:  8,
		6:  10,
		7:  12,
		8:  14,
		9:  16,
		10: 18,
		11: 20,
		12: 25,
		13: 30,
		14: 35,
		15: 40,
		16: 45,
		17: 50,
		18: 60,
		19: 80,
		20: 100,
		21: 120,
		22: 140,
		23: 160,
		24: 180,
		25: 200,
		26: 220,
		27: 240,
		28: 260,
		29: 300,
		30: 350,
		31: 400,
		32: 450,
		33: 500,
		34: 550,
		35: 600,
		36: 600,
	},
	GAME_SOLT_6: {
		5:  8,
		6:  10,
		7:  12,
		8:  14,
		9:  16,
		10: 18,
		11: 20,
		12: 25,
		13: 30,
		14: 35,
		15: 40,
		16: 45,
		17: 50,
		18: 60,
		19: 80,
		20: 100,
		21: 120,
		22: 140,
		23: 160,
		24: 180,
		25: 220,
		26: 250,
		27: 280,
		28: 320,
		29: 360,
		30: 400,
		31: 450,
		32: 500,
		33: 550,
		34: 600,
		35: 800,
		36: 800,
	},
	GAME_SOLT_7: {
		5:  10,
		6:  12,
		7:  15,
		8:  20,
		9:  25,
		10: 30,
		11: 35,
		12: 40,
		13: 45,
		14: 50,
		15: 60,
		16: 70,
		17: 80,
		18: 100,
		19: 120,
		20: 150,
		21: 200,
		22: 250,
		23: 300,
		24: 350,
		25: 400,
		26: 450,
		27: 500,
		28: 600,
		29: 800,
		30: 1000,
		31: 1200,
		32: 1500,
		33: 1800,
		34: 2000,
		35: 2500,
		36: 2500,
	},
	GAME_SOLT_8: {
		5:  10,
		6:  12,
		7:  15,
		8:  20,
		9:  25,
		10: 30,
		11: 35,
		12: 40,
		13: 45,
		14: 50,
		15: 60,
		16: 70,
		17: 80,
		18: 100,
		19: 120,
		20: 150,
		21: 200,
		22: 250,
		23: 300,
		24: 350,
		25: 400,
		26: 450,
		27: 500,
		28: 600,
		29: 800,
		30: 1000,
		31: 1200,
		32: 1500,
		33: 1800,
		34: 2000,
		35: 2500,
		36: 2500,
	},
}

type Prize struct {
	slot    int
	count   int
	mult    int
	addmult int
	value   int64
	idxs    []int
}
type SlotInfo struct {
	slot int
	idx  int
}

type RoundInfo struct {
	cur       [BUFF_SIZE]int //当前
	final     [BUFF_SIZE]int //最终
	remove    []SlotInfo     //移除
	add       []SlotInfo     //添加
	prices    []Prize        //中奖
	special   []int          //炸弹，火箭，地雷位置的索引
	curmult   []int          //当前乘数区
	finalmult []int          //最终乘数区
}

var rd = rand.New(rand.NewSource(time.Now().UnixNano()))

func randBuffer(out *[BUFF_SIZE]int) {
	var Buffers = []int{
		GAME_SOLT_1,
		GAME_SOLT_2,
		GAME_SOLT_3,
		GAME_SOLT_4,
		GAME_SOLT_5,
		GAME_SOLT_6,
		GAME_SOLT_7,
		GAME_SOLT_8,
		GAME_SOLT_WILD,
		GAME_SOLT_SCATTER,
	}
	for col := 0; col < COL_DEF; col++ {
		for row := 0; row < ROW_DEF; row++ {
			idx := row*COL_DEF + col
			out[idx] = Buffers[rd.Intn(len(Buffers))]
		}
	}
}

func randFreeBuffer(out *[BUFF_SIZE]int, mult *[BUFF_SIZE]int) (cp []int) {
	var Buffers = []int{
		GAME_SOLT_1,
		GAME_SOLT_2,
		GAME_SOLT_3,
		GAME_SOLT_4,
		GAME_SOLT_5,
		GAME_SOLT_6,
		GAME_SOLT_7,
		GAME_SOLT_8,
		GAME_SOLT_WILD,
	}
	cp = make([]int, BUFF_SIZE)
	for col := 0; col < COL_DEF; col++ {
		for row := 0; row < ROW_DEF; row++ {
			idx := row*COL_DEF + col
			if mult[idx] > 1 {
				out[idx] = GAME_SOLT_WILD
			} else {
				out[idx] = Buffers[rd.Intn(len(Buffers))]
			}
			cp[idx] = mult[idx]
		}
	}
	return
}

func printfRoundInfo(i int, round *RoundInfo, bfree bool, fcout int) {
	strSlot := [GAME_SOLT_MAX]string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"w",
		"a",
		"b",
		"c",
		"d",
	}
	sb := strings.Builder{}
	bCMult := false
	bFMult := false
	if bfree {
		sb.WriteString(fmt.Sprintf("\nFREE:%d.%d w:WILD,a:BOMB,b:ROCKET,c:MINE,d:SCATTER", fcout, i))
		for i := 0; i < len(round.curmult); i++ {
			if round.curmult[i] > 1 {
				bCMult = true
				break
			}
		}
		for i := 0; i < len(round.curmult); i++ {
			if round.finalmult[i] > 1 {
				bFMult = true
				break
			}
		}
	} else {
		sb.WriteString(fmt.Sprintf("\n====<%d> w:WILD,a:BOMB,b:ROCKET,c:MINE,d:SCATTER", i))
	}
	if len(round.special) > 0 {
		sb.WriteString(" <!BOMB OR ROCKET OR MINE!>\n")
	} else {
		sb.WriteString("\n")
	}
	tmp := round.cur
	for _, v := range round.remove {
		tmp[v.idx] = 0
	}
	for row := 0; row < ROW_DEF; row++ {
		for col := 0; col < COL_DEF; col++ {
			idx := row*COL_DEF + col
			if bCMult {
				sb.WriteString(fmt.Sprintf("[%s,%d]", strSlot[round.cur[idx]], round.curmult[idx]))
			} else {
				sb.WriteString(strSlot[round.cur[idx]])
			}
			if col < ROW_DEF-1 {
				sb.WriteString(",")
			}
		}
		if row == 2 {
			sb.WriteString(" >>>> ")
		} else {
			sb.WriteString("      ")
		}
		for col := 0; col < COL_DEF; col++ {
			idx := row*COL_DEF + col
			if tmp[idx] == 0 {
				sb.WriteString("*")
			} else {
				sb.WriteString(strSlot[tmp[idx]])
			}
			if col < ROW_DEF-1 {
				sb.WriteString(",")
			}
		}
		if row == 2 {
			sb.WriteString(" >>>> ")
		} else {
			sb.WriteString("      ")
		}
		for col := 0; col < COL_DEF; col++ {
			idx := row*COL_DEF + col
			if bFMult {
				sb.WriteString(fmt.Sprintf("[%s,%d]", strSlot[round.final[idx]], round.finalmult[idx]))
			} else {
				sb.WriteString(strSlot[round.final[idx]])
			}
			if col < ROW_DEF-1 {
				sb.WriteString(",")
			}
		}
		if row < COL_DEF-1 {
			sb.WriteString("\n")
		}
	}
	sb.WriteString("\n")
	if len(round.prices) > 0 {
		//sb.WriteString("price info\n")
	}
	for _, price := range round.prices {
		sb.WriteString(fmt.Sprintf("bet:100 slot:%d count:%d mult:%d addmult:%d wins:%d\n", price.slot, price.count, price.mult, price.addmult, price.value))
	}
	fmt.Print(sb.String())
}

func isValid(x, y int) bool {
	return x >= 0 && x < ROW_DEF && y >= 0 && y < COL_DEF
}

func dfs(buff *[BUFF_SIZE]int, row, col int, target int) []int {
	idx := row*COL_DEF + col
	stack := []int{idx}
	group := []int{idx}
	visited := [BUFF_SIZE]bool{}
	visited[idx] = true
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currX := curr / ROW_DEF
		currY := curr % COL_DEF
		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range directions {
			nextX := currX + dir[0]
			nextY := currY + dir[1]
			next := nextX*COL_DEF + nextY
			if isValid(nextX, nextY) && !visited[next] {
				if buff[next] == target || buff[next] == GAME_SOLT_WILD {
					visited[next] = true
					stack = append(stack, next)
					group = append(group, next)
				}
			}
		}
	}
	return group
}

func findSameElements(buff *[BUFF_SIZE]int) [][]int {
	var result [][]int
	visited := [BUFF_SIZE]bool{}
	for col := 0; col < COL_DEF; col++ {
		for row := 0; row < ROW_DEF; row++ {
			idx := row*COL_DEF + col
			switch buff[idx] {
			case GAME_SOLT_WILD:
				fallthrough
			case GAME_SOLT_BOMB:
				fallthrough
			case GAME_SOLT_ROCKET:
				fallthrough
			case GAME_SOLT_MINE:
				fallthrough
			case GAME_SOLT_SCATTER:
				continue
			}
			if !visited[idx] {
				group := dfs(buff, row, col, buff[idx])
				if len(group) >= 5 {
					result = append(result, group)
				}
				for _, v := range group {
					visited[v] = true
				}
			}
		}
	}
	return result
}

func getMult(slot int, count int) int {
	if m, ok := Mult[slot]; ok {
		if val, b := m[count]; b {
			return val
		}
	}
	return 0
}

func calcSocre(betScore int, info *RoundInfo, mult *[BUFF_SIZE]int) int64 {
	reset := map[int]int{}
	elements := findSameElements(&info.cur)
	for _, element := range elements {
		addMult := 0
		solt := GAME_SOLT_EMPTY
		for i := 0; i < len(element); i++ {
			idx := element[i]
			if info.cur[idx] != GAME_SOLT_WILD {
				solt = info.cur[idx]
			}
			info.remove = append(info.remove, SlotInfo{
				slot: info.cur[idx],
				idx:  idx,
			})
			if mult != nil && mult[element[i]] > 1 {
				addMult += mult[idx]
				reset[idx] = idx
			}
		}
		if addMult == 0 {
			addMult = 1
		}
		count := len(element)
		mult := getMult(solt, count)
		info.prices = append(info.prices, Prize{
			slot:    solt,
			count:   count,
			mult:    mult,
			addmult: addMult,
			value:   int64(betScore) * int64(mult) * int64(addMult),
			idxs:    element,
		})
	}
	//倍数用完就没了
	for k := range reset {
		mult[k] = 1
		info.finalmult[k] = 1
	}
	if len(elements) > 0 {
		return 1
	}
	return 0
}

func randDropBuffer(round *RoundInfo, mult *[BUFF_SIZE]int) {
	//赋值
	round.final = round.cur
	//原始图案
	buff := round.final
	//消除
	for _, v := range round.remove {
		round.final[v.idx] = GAME_SOLT_EMPTY
	}
	//根据消除个数生成炸弹、火箭或者地雷
	//消除大于6个可以生成炸弹
	//消除大于9个可以生成火箭
	//消除大于12个可以生成地雷
	for i := 0; i < len(round.prices); i++ {
		slot := GAME_SOLT_EMPTY
		if round.prices[i].count >= 6 && round.prices[i].count < 9 {
			slot = GAME_SOLT_BOMB
		} else if round.prices[i].count >= 9 && round.prices[i].count < 12 {
			slot = GAME_SOLT_ROCKET
		} else if round.prices[i].count >= 12 {
			slot = GAME_SOLT_MINE
		}
		if slot != GAME_SOLT_EMPTY {
			//WILD位置不能生成图案
			tmp := []int{}
			//过滤WILD的位置
			for _, idx := range round.prices[i].idxs {
				if buff[idx] != GAME_SOLT_WILD {
					tmp = append(tmp, idx)
				}
			}
			idx := tmp[rd.Intn(len(tmp))]
			round.final[idx] = slot
			round.add = append(round.add, SlotInfo{
				slot: slot,
				idx:  idx,
			})
		}
	}
	//掉落添加
	round.add = drop(&round.final, mult)
}

func drop(out *[BUFF_SIZE]int, mult *[BUFF_SIZE]int) (add []SlotInfo) {
	var Buffers = []int{
		GAME_SOLT_1,
		GAME_SOLT_2,
		GAME_SOLT_3,
		GAME_SOLT_4,
		GAME_SOLT_5,
		GAME_SOLT_6,
		GAME_SOLT_7,
		GAME_SOLT_8,
		GAME_SOLT_WILD,
	}
	// out[0] = 4
	// out[6] = 13
	// out[12] = 13
	// out[18] = 1
	// out[24] = 0
	// out[30] = 0

	for col := 0; col < COL_DEF; col++ {
		// sb := strings.Builder{}
		// for row := 0; row < ROW_DEF; row++ {
		// 	idx := row*COL_DEF + col
		// 	sb.WriteString(fmt.Sprintf("%02d,", out[idx]))
		// 	if row == ROW_DEF-1 {
		// 		sb.WriteString("\n")
		// 	}
		// }
		// fmt.Print("pre:" + sb.String())
		//下坠
		for i := ROW_DEF - 1; i >= 0; i-- {
			idxi := i*COL_DEF + col
			if out[idxi] == GAME_SOLT_EMPTY {
				swap := -1
				for j := i - 1; j >= 0; j-- {
					idxj := j*COL_DEF + col
					//特殊元素不能下坠
					switch out[idxj] {
					case GAME_SOLT_EMPTY:
						fallthrough
					case GAME_SOLT_BOMB:
						fallthrough
					case GAME_SOLT_ROCKET:
						fallthrough
					case GAME_SOLT_MINE:
						continue
					}
					//有乘数的
					if mult != nil && mult[idxj] > 1 {
						continue
					}
					swap = idxj
					break
				}
				if swap > -1 {
					out[idxi] = out[swap]
					out[swap] = GAME_SOLT_EMPTY
				}
			}
		}
		// sb = strings.Builder{}
		// for row := 0; row < ROW_DEF; row++ {
		// 	idx := row*COL_DEF + col
		// 	sb.WriteString(fmt.Sprintf("%02d,", out[idx]))
		// 	if row == ROW_DEF-1 {
		// 		sb.WriteString("\n")
		// 	}
		// }
		// fmt.Print("drop:" + sb.String())
		//填充
		for row := ROW_DEF - 1; row >= 0; row-- {
			idx := row*COL_DEF + col
			if out[idx] == GAME_SOLT_EMPTY {
				slotType := Buffers[rd.Intn(len(Buffers))]
				add = append(add, SlotInfo{
					slot: slotType,
					idx:  idx,
				})
				out[idx] = slotType
			}
		}
		// sb = strings.Builder{}
		// for row := 0; row < ROW_DEF; row++ {
		// 	idx := row*COL_DEF + col
		// 	sb.WriteString(fmt.Sprintf("%02d,", out[idx]))
		// 	if row == ROW_DEF-1 {
		// 		sb.WriteString("\n")
		// 	}
		// }
		// fmt.Print("last:" + sb.String())
	}
	// 	//列缓存
	// 	colarrs := []int{}
	// 	//三个特殊元素不可以掉落
	// 	for row := 0; row < ROW_DEF; row++ {
	// 		idx := row*COL_DEF + col
	// 		if out[idx] != GAME_SOLT_EMPTY {
	// 			colarrs = append(colarrs, out[idx])
	// 		}
	// 	}
	// 	dropCount := ROW_DEF - len(colarrs)
	// 	for i := 0; i < dropCount; i++ {
	// 		colarrs = append([]int{GAME_SOLT_EMPTY}, colarrs...)
	// 	}
	// 	if dropCount > 0 {
	// 		for row := 0; row < ROW_DEF; row++ {
	// 			if colarrs[row] == GAME_SOLT_EMPTY {
	// 				slotType := Buffers[rd.Intn(len(Buffers))]
	// 				colarrs[row] = slotType
	// 				add = append(add, SlotInfo{
	// 					slot: slotType,
	// 					idx:  row*COL_DEF + col,
	// 				})
	// 			}
	// 		}
	// 		//用最新的图
	// 		for row, v := range colarrs {
	// 			idx := row*COL_DEF + col
	// 			out[idx] = v
	// 		}
	// 	}
	// }
	return
}

func checkSpecialBuff(round *RoundInfo) {
	//赋值
	round.final = round.cur
	//检查
	for i := 0; i < BUFF_SIZE; i++ {
		if round.final[i] == GAME_SOLT_BOMB || round.final[i] == GAME_SOLT_ROCKET || round.final[i] == GAME_SOLT_MINE {
			round.special = append(round.special, i)
		}
	}
}

func randDropSpecialBuffer(round *RoundInfo, mult *[BUFF_SIZE]int) {
	remove := [BUFF_SIZE]bool{}
	//图案消除
	for i := 0; i < len(round.special); i++ {
		idx := round.special[i]
		switch round.final[idx] {
		case GAME_SOLT_BOMB:
			bombRemove(idx, &round.final, &remove, mult)
		case GAME_SOLT_ROCKET:
			rocketRemove(idx, &round.final, &remove, mult)
		case GAME_SOLT_MINE:
			mineRemove(idx, &round.final, &remove)
		}
	}
	for i := 0; i < BUFF_SIZE; i++ {
		//记录消除
		if remove[i] {
			round.remove = append(round.remove, SlotInfo{
				slot: round.final[i],
				idx:  i,
			})
			round.final[i] = GAME_SOLT_EMPTY
		}
		//免费的乘数区
		if mult != nil {
			round.finalmult[i] = mult[i]
		}
	}
	//掉落添加
	round.add = drop(&round.final, mult)
}

func bombRemove(idx int, buff *[BUFF_SIZE]int, remove *[BUFF_SIZE]bool, mult *[BUFF_SIZE]int) (c int) {
	remove[idx] = true
	//方向
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}}
	currX := idx / ROW_DEF
	currY := idx % COL_DEF
	//消除3*3
	for _, dir := range directions {
		nextX := currX + dir[0]
		nextY := currY + dir[1]
		if isValid(nextX, nextY) {
			next := nextX*COL_DEF + nextY
			if buff[next] < GAME_SOLT_WILD {
				remove[next] = true
			} else if buff[next] == GAME_SOLT_WILD {
				if mult != nil && mult[next] > 0 {
					mult[next]++
					c++
				}
			}
		}
	}
	// sb := strings.Builder{}
	// for row := 0; row < ROW_DEF; row++ {
	// 	for col := 0; col < COL_DEF; col++ {
	// 		idx := row*COL_DEF + col
	// 		if remove[idx] {
	// 			sb.WriteString("*")
	// 		} else {
	// 			sb.WriteString(fmt.Sprintf("%x", buff[idx]))
	// 		}
	// 		if col < ROW_DEF-1 {
	// 			sb.WriteString(",")
	// 		}
	// 	}
	// 	sb.WriteString("\n")
	// }
	// fmt.Print(sb.String())
	return
}

func rocketRemove(idx int, buff *[BUFF_SIZE]int, remove *[BUFF_SIZE]bool, mult *[BUFF_SIZE]int) (c int) {
	remove[idx] = true

	currX := idx / ROW_DEF
	currY := idx % COL_DEF

	//消除横向
	for i := 0; i < COL_DEF; i++ {
		next := currX*COL_DEF + i
		if buff[next] < GAME_SOLT_WILD {
			remove[next] = true
		} else if buff[next] == GAME_SOLT_WILD {
			if mult != nil && mult[next] > 0 {
				mult[next]++
				c++
			}
		}
	}
	//消除纵向
	for i := 0; i < ROW_DEF; i++ {
		next := i*COL_DEF + currY
		if buff[next] < GAME_SOLT_WILD {
			remove[next] = true
		} else if buff[next] == GAME_SOLT_WILD {
			if mult != nil && mult[next] > 0 {
				mult[next]++
				c++
			}
		}
	}

	// sb := strings.Builder{}
	// for row := 0; row < ROW_DEF; row++ {
	// 	for col := 0; col < COL_DEF; col++ {
	// 		idx := row*COL_DEF + col
	// 		if remove[idx] {
	// 			sb.WriteString("*")
	// 		} else {
	// 			sb.WriteString(fmt.Sprintf("%x", buff[idx]))
	// 		}
	// 		if col < ROW_DEF-1 {
	// 			sb.WriteString(",")
	// 		}
	// 	}
	// 	sb.WriteString("\n")
	// }
	// fmt.Print(sb.String())
	return
}

func mineRemove(idx int, buff *[BUFF_SIZE]int, remove *[BUFF_SIZE]bool) {
	remove[idx] = true
	//消除全部
	for i := 0; i < BUFF_SIZE; i++ {
		if buff[i] < GAME_SOLT_WILD {
			remove[i] = true
		}
	}
	// sb := strings.Builder{}
	// for row := 0; row < ROW_DEF; row++ {
	// 	for col := 0; col < COL_DEF; col++ {
	// 		idx := row*COL_DEF + col
	// 		if remove[idx] {
	// 			sb.WriteString("*")
	// 		} else {
	// 			sb.WriteString(fmt.Sprintf("%x", buff[idx]))
	// 		}
	// 		if col < ROW_DEF-1 {
	// 			sb.WriteString(",")
	// 		}
	// 	}
	// 	sb.WriteString("\n")
	// }
	// fmt.Print(sb.String())
}

func checkTrigerFreeCount(buff *[BUFF_SIZE]int) int {
	count := 0
	for i := 0; i < BUFF_SIZE; i++ {
		if buff[i] == GAME_SOLT_SCATTER {
			count++
		}
	}
	if count >= 3 {
		return 10
	}
	return 0
}
