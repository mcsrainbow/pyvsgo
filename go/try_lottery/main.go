// 包名为 main，表示这是一个可执行程序
package main

import (
	"fmt"       // fmt 包提供了格式化输出的函数
	"math/rand" // math/rand 包提供了随机数生成的功能
	"os"        // os 包提供了操作系统相关的功能
	"sort"      // sort 包提供了排序算法的实现
	"time"      // time 包提供了时间相关的功能
)

// Lottery 结构体，用于存储用户选择的球、生成的球和随机数源等信息
type Lottery struct {
	userBalls  []int      // 用户选择的球，长度为 TotalBallCount
	redBalls   []int      // 红球，长度为 RedBallCount
	blueBall   int        // 蓝球
	randSource *rand.Rand // 随机数源
}

// 常量定义：红球最大值、蓝球最大值、红球数量和总球数量
const (
	RedBallMax     = 33 // 红球的最大值为 33
	BlueBallMax    = 16 // 蓝球的最大值为 16
	RedBallCount   = 6  // 红球的数量为 6
	TotalBallCount = 7  // 总球数量为 7
)

// ValidateUserBalls 方法：验证用户选择的球是否有效
func (lottery *Lottery) ValidateUserBalls() {
	// 检查用户选择的球的长度是否正确
	if len(lottery.userBalls) != TotalBallCount {
		fmt.Println("NOTICE: 你必须提供 exactly 7 个球。") // 输出错误信息
		os.Exit(1)                                 // 退出程序
	}
	// 检查每个球是否为正整数
	for _, ball := range lottery.userBalls {
		if ball <= 0 {
			fmt.Printf("NOTICE: 球 %d 不是有效的正整数。\n", ball) // 输出错误信息
			os.Exit(1)                                   // 退出程序
		}
	}
	// 检查红球是否在有效范围内
	redBalls := lottery.userBalls[:RedBallCount]    // 获取红球
	blueBall := lottery.userBalls[TotalBallCount-1] // 获取蓝球
	for _, ball := range redBalls {
		if ball > RedBallMax {
			fmt.Println("NOTICE: 一个红球的值超过了最大允许值 (33)。") // 输出错误信息
			os.Exit(1)                                  // 退出程序
		}
	}
	// 检查蓝球是否在有效范围内
	if blueBall > BlueBallMax {
		fmt.Printf("NOTICE: 蓝球 %d 的值超过了最大允许值 (16)。\n", blueBall) // 输出错误信息
		os.Exit(1)                                               // 退出程序
	}
}

// GenerateLotteryBalls 方法：生成所有的彩票球
func (lottery *Lottery) GenerateLotteryBalls() []int {
	// 生成一个从 1 到 RedBallMax 的随机排列
	numbers := lottery.randSource.Perm(RedBallMax)
	for i := range numbers {
		numbers[i] += 1 // 确保数字在 1 到 RedBallMax 范围内
	}
	lottery.redBalls = numbers[:RedBallCount] // 获取红球
	sort.Ints(lottery.redBalls)               // 排序红球

	// 生成蓝球
	lottery.blueBall = lottery.randSource.Intn(BlueBallMax) + 1 // 生成一个从 1 到 BlueBallMax 的随机数

	return append(lottery.redBalls, lottery.blueBall) // 返回所有的彩票球
}

// CompareBalls 函数：比较两个球切片是否相等
func CompareBalls(balls1, balls2 []int) bool {
	// 检查长度是否相同
	if len(balls1) != len(balls2) {
		return false // 长度不同，返回 false
	}
	// 检查每个球是否相等
	for i, ball := range balls1 {
		if ball != balls2[i] {
			return false // 球不相等，返回 false
		}
	}
	return true // 所有球都相等，返回 true
}

func main() {
	// 创建一个新的随机数源
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 用户选择的球：1, 5, 10, 15, 16, 26, 9
	userBalls := []int{1, 5, 10, 15, 16, 26, 9}
	// 创建一个新的 Lottery 对象
	lottery := Lottery{
		userBalls:  userBalls,
		randSource: randSource,
	}
	// 验证用户选择的球是否有效
	lottery.ValidateUserBalls()

	attempts := 1 // 尝试次数

	for {
		// 生成所有的彩票球
		generatedBalls := lottery.GenerateLotteryBalls()
		// 检查生成的球是否与用户选择的球相等
		if CompareBalls(generatedBalls, userBalls) {
			fmt.Printf("\nNOTICE: 你终于中了 500W RMB，买了 %d 张票！\n", attempts) // 输出中奖信息
			break                                                        // 中奖后退出循环
		}
		fmt.Printf("\rINFO: 你已经尝试了 %d 次。", attempts) // 输出尝试次数
		attempts++                                   // 尝试次数加 1
	}
}
