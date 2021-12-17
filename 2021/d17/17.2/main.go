package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

const Y_RANGE_MIN = -126
const Y_RANGE_MAX = -69
const X_RANGE_MIN = 217
const X_RANGE_MAX = 240

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve() {
	minXVel := int(math.Sqrt(X_RANGE_MIN*2+0.25)-0.5) + 1

	count := 0
	for xVel := minXVel; xVel <= X_RANGE_MAX; xVel++ {

		for yVel := -1000; yVel < 1000; yVel++ {
			if Launch(xVel, yVel) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func Launch(xVel, yVel int) bool {
	var xPos, yPos int

	if yVel > 0 {
		xVelTmp := xVel
		xVelTmp = xVelTmp - yVel
		if xVelTmp < 0 {
			xVelTmp = 0
		}

		xPos = ((xVel-xVelTmp+1)*(xVel+xVelTmp))/2 - xVelTmp
		xVel = xVelTmp

		yPos = (yVel*yVel + yVel) / 2
		yVel = 0
	}

	for {
		yPos, yVel = stepY(yPos, yVel)
		xPos, xVel = stepX(xPos, xVel)

		if isMiss(xPos, yPos, xVel) {
			break
		} else if isHit(xPos, yPos) {
			return true
		}
	}

	return false
}

func isMiss(xPos, yPos, xVel int) bool {
	if yPos < Y_RANGE_MIN || xPos > X_RANGE_MAX || (xVel == 0 && xPos < X_RANGE_MIN) {
		return true
	}

	return false
}

func isHit(xPos, yPos int) bool {
	if (yPos <= Y_RANGE_MAX && yPos >= Y_RANGE_MIN) && (xPos <= X_RANGE_MAX && xPos >= X_RANGE_MIN) {
		return true
	}

	return false
}

func stepY(yPos, yVel int) (int, int) {
	if yVel >= 0 {
		return yPos - yVel, yVel + 1
	} else {
		return yPos + yVel, yVel - 1
	}
}

func stepX(xPos, xVel int) (int, int) {
	if xVel <= 0 {
		return xPos, 0
	} else {
		return xPos + xVel, xVel - 1
	}
}
