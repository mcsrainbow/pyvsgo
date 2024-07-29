package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

// Lottery struct to hold user balls, generated balls, and constants
type Lottery struct {
	userBalls []int
	redBalls  []int
	blueBall  int
}

// Constants for the maximum values and the count of balls
const (
	RedBallMax     = 33
	BlueBallMax    = 16
	RedBallCount   = 6
	TotalBallCount = 7
)

// ValidateUserBalls validates the user provided balls
func (lottery *Lottery) ValidateUserBalls() {
	if len(lottery.userBalls) != TotalBallCount {
		fmt.Println("NOTICE: You must provide exactly 7 balls.")
		os.Exit(1)
	}
	for _, ball := range lottery.userBalls {
		if ball <= 0 {
			fmt.Printf("NOTICE: Ball %d is not a valid positive integer.\n", ball)
			os.Exit(1)
		}
	}
	redBalls := lottery.userBalls[:RedBallCount]
	blueBall := lottery.userBalls[TotalBallCount-1]
	for _, ball := range redBalls {
		if ball > RedBallMax {
			fmt.Println("NOTICE: One of the red balls is greater than the maximum allowed (33).")
			os.Exit(1)
		}
	}
	if blueBall > BlueBallMax {
		fmt.Printf("NOTICE: Blue ball %d is greater than the maximum allowed (16).\n", blueBall)
		os.Exit(1)
	}
}

// GenerateLotteryBalls generates all lottery balls
func (lottery *Lottery) GenerateLotteryBalls() []int {
	// Generate a permutation of numbers from 1 to RedBallMax
	numbers := rand.Perm(RedBallMax)
	for i := range numbers {
		numbers[i] += 1 // Ensure the numbers are in the range 1 to RedBallMax
	}
	lottery.redBalls = numbers[:RedBallCount]
	sort.Ints(lottery.redBalls)

	// Generate blue ball
	lottery.blueBall = rand.Intn(BlueBallMax) + 1

	return append(lottery.redBalls, lottery.blueBall)
}

// CompareBalls compares two slices of balls
func CompareBalls(balls1, balls2 []int) bool {
	if len(balls1) != len(balls2) {
		return false
	}
	for i, ball := range balls1 {
		if ball != balls2[i] {
			return false
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	userBalls := []int{1, 5, 10, 15, 16, 26, 9} // 用户选择的球
	lottery := Lottery{userBalls: userBalls}
	lottery.ValidateUserBalls()

	attempts := 1

	for {
		generatedBalls := lottery.GenerateLotteryBalls()
		if CompareBalls(generatedBalls, userBalls) {
			fmt.Printf("\nNOTICE: You finally got 500W RMB after buying %d tickets!\n", attempts)
			break
		}
		fmt.Printf("\rINFO: You have tried %d times.", attempts)
		attempts++
	}
}
