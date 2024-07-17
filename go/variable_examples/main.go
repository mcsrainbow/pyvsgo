package main

import (
    "fmt"
)

// 全局常量
const globalConst = "I am a global constant"

// 对外部包可见的全局常量, 首字母大写
const GlobalConst = "I am a global constant visible to external packages"

// 全局变量
var globalVar = "I am a global variable"

// 对外部包可见的全局变量, 首字母大写
var GlobalVar = "I am a global variables visible to external packages"

// 静态变量, 对比 Python 类变量
var staticVar = 10

// 函数, 增加类变量的值
func incrementStaticVar() {
    staticVar++
}

// 定义结构体类型, 对比 Python Class 类
type MyClass struct {
    // 实例变量
    instanceVar int
}

// 方法, 增加实例变量的值
func (obj *MyClass) incrementInstanceVar() {
    obj.instanceVar++
}

func myFunction() {
    // 局部变量
    localVar := "myFunction() localVar: I am a local variable"
    fmt.Println(localVar)
}

func main() {
    // 使用示例
    fmt.Println("globalConst:", globalConst)
    fmt.Println("GlobalConst:", GlobalConst)
    fmt.Println("globalVar:", globalVar)
    fmt.Println("GlobalVar:", GlobalVar)

    // 动态变量, Go 中通过空接口实现
    var dynamicVar interface{}
    dynamicVar = 5
    fmt.Println("dynamicVar:", dynamicVar)
    dynamicVar = "Now I'm a string"
    fmt.Println("dynamicVar:", dynamicVar)

    // 切片 (List in Python)
    mySlice := []int{1, 2, 3, 4, 5}

    // 映射 (Dict in Python)
    myMap := map[string]interface{}{"name": "John", "age": 30}

    // 字符串
    myString := "Hello, World!"

    // 使用示例
    myFunction()
    fmt.Println("mySlice:", mySlice)
    fmt.Println("myMap:", myMap)
    fmt.Println("myString:", myString)

    fmt.Println("staticVar:", staticVar)
    // 增加类变量
    incrementStaticVar()
    fmt.Println("incrementStaticVar() staticVar:", staticVar)

    // 创建类实例
    obj := MyClass{instanceVar: 1}

    fmt.Println("obj.instanceVar:", obj.instanceVar)
    // 增加实例变量
    obj.incrementInstanceVar()
    fmt.Println("obj.incrementInstanceVar() obj.instanceVar:", obj.instanceVar)
}
