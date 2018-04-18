package main

import (
	"math/rand"
	"time"
)

func main() {

	time := time.Now().UnixNano()
	rand.Seed(time)

	dice := rand.Intn(6) + 1
	println(dice, uranai(dice))
}

func uranai(i int) string {
	switch i {
	case 1:
		{
			return "凶"
		}
	case 2, 3:
		{
			return "吉"
		}
	case 4, 5:
		{
			return "中吉"
		}
	case 6:
		{
			return "大吉"
		}
	}
	return "？？？"
}
