package main

import (
	"fmt"
	"log"
	"time"
)

const Y_RANGE_MIN = -126
const Y_RANGE_MAX = -69

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve() {
	for i := 1000; i > 0; i-- {
		yMax := (i*i + i) / 2

		yVel := 0
		yPos := yMax
		for {
			yPos = yPos - yVel
			if yPos < Y_RANGE_MIN {
				break
			} else if yPos <= Y_RANGE_MAX && yPos >= Y_RANGE_MIN {
				fmt.Println(yMax)
				return
			}

			yVel = yVel + 1
		}
	}
}
