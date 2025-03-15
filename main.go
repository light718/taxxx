package main

import "time"

func do(betScore int, mult *[BUFF_SIZE]int) (rounds []RoundInfo) {
	for {
		round := RoundInfo{}
		if len(rounds) <= 0 {
			if mult == nil {
				randBuffer(&round.cur)
			} else {
				round.curmult = randFreeBuffer(&round.cur, mult)
				round.finalmult = append(round.finalmult, round.curmult...)
			}
		} else {
			plast := &rounds[len(rounds)-1]
			round.cur = plast.final
			round.curmult = append(round.curmult, plast.finalmult...)
			round.finalmult = append(round.finalmult, round.curmult...)
		}
		calcSocre(betScore, &round, mult)
		if len(round.remove) > 0 {
			randDropBuffer(&round, mult)
		} else {
			checkSpecialBuff(&round)
			if len(round.special) <= 0 {
				if len(rounds) <= 0 {
					rounds = append(rounds, round)
				}
				break
			}
			randDropSpecialBuffer(&round, mult)
		}
		rounds = append(rounds, round)
	}
	return
}

func main() {
	betScore := 100
	for {
		rounds := do(betScore, nil)
		//打印
		for i := 0; i < len(rounds); i++ {
			printfRoundInfo(i+1, &rounds[i], false, 0)
		}
		//检查免费
		mult := [BUFF_SIZE]int{}
		freeCount := checkTrigerFreeCount(&rounds[len(rounds)-1].final)
		if freeCount > 0 {
			for i := 0; i < BUFF_SIZE; i++ {
				//初始化倍率
				mult[i] = 1
			}
		}
		fcount := 0
		for freeCount > 0 {
			freeRounds := do(betScore, &mult)
			for i := 0; i < len(freeRounds); i++ {
				printfRoundInfo(i+1, &freeRounds[i], true, fcount+1)
			}
			addCount := checkTrigerFreeCount(&freeRounds[len(freeRounds)-1].final)
			freeCount += addCount
			freeCount--
			fcount++
		}
		time.Sleep(time.Millisecond * time.Duration(1000))
	}
}
