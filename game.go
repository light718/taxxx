package main

import "time"

func main() {
	for {
		rounds := []RoundInfo{}
		for {
			round := RoundInfo{}
			if len(rounds) <= 0 {
				randBuffer(&round.cur)
			} else {
				round.cur = rounds[len(rounds)-1].final
			}
			calcSocre(100, &round)
			if len(round.remove) > 0 {
				randDropBuffer(&round)
			} else {
				checkSpecialBuff(&round)
				if len(round.special) <= 0 {
					if len(rounds) <= 0 {
						rounds = append(rounds, round)
					}
					break
				}
				randDropSpecialBuffer(&round)
			}
			rounds = append(rounds, round)
		}
		for i := 0; i < len(rounds); i++ {
			printfRoundInfo(i+1, &rounds[i])
		}
		time.Sleep(time.Millisecond * time.Duration(100))
	}
}
