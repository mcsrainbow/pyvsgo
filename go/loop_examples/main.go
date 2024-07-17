package main

import (
	"fmt"
)

// Go 中使用 switch-case 语句
func matchCaseExample(value int) string {
	switch value {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	default:
		return "Other"
	}
}

func main() {
	// 使用 for 循环遍历 slice
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Using 'for' loop to iterate through slice:")
	for _, item := range mySlice {
		fmt.Println(item)
	}

	// 使用 while 循环遍历 slice (在 Go 中没有 while 循环，使用 for 代替)
	fmt.Println("\nUsing 'while' loop to iterate through slice:")
	index := 0
	for index < len(mySlice) {
		fmt.Println(mySlice[index])
		index++
	}

	// 使用 for 循环遍历 map
	myMap := map[string]interface{}{
		"name":      "John",
		"age":       30,
		"isStudent": false,
		"scores":    []int{90, 85, 92},
	}
	fmt.Println("\nUsing 'for' loop to iterate through map:")
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// 使用 while 循环遍历 map (在 Go 中没有 while 循环，使用 for 代替)
	fmt.Println("\nUsing 'while' loop to iterate through map:")
	keys := make([]string, 0, len(myMap))
	for key := range myMap {
		keys = append(keys, key)
	}
	index = 0
	for index < len(keys) {
		key := keys[index]
		fmt.Printf("Key: %s, Value: %d\n", key, myMap[key])
		index++
	}

	// Go 中没有类似 Python 3.10 的 match-case 语句, 使用 switch 代替
	fmt.Println("\nUsing 'switch' statement:")
	fmt.Println(matchCaseExample(1))
	fmt.Println(matchCaseExample(4))

	// 死循环示例
	fmt.Println("\nInfinite loop example:")
	for {
		fmt.Println("This is an infinite loop. Use 'break' to exit.")
		break // 添加 break 以防止实际运行时进入死循环
	}

	// 使用 break 退出循环的示例
	fmt.Println("\nUsing 'break' to exit the loop at 'number 5':")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 等于5时退出循环
		}
		fmt.Println(i)
	}

	// 使用 continue 在循环中的示例 (在 Go 中没有 pass 关键字, 直接使用 continue)
	fmt.Println("\nUsing 'continue' in the loop at 'even numbers':")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 偶数时什么都不做
		} else {
			fmt.Println(i) // 奇数时打印
		}
	}
}

