package main

import (
	"fmt"
)

// OrderedMap 是有序字典的结构体
// keys 切片保持插入顺序
// values 映射存储键值对
type OrderedMap struct {
	keys   []string
	values map[string]interface{}
}

// NewOrderedMap 创建一个新的 OrderedMap
func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		keys:   []string{},
		values: make(map[string]interface{}),
	}
}

// Set 插入或更新键值对
func (om *OrderedMap) Set(key string, value interface{}) {
	if _, exists := om.values[key]; !exists {
		om.keys = append(om.keys, key)
	}
	om.values[key] = value
}

// Get 获取键的值
func (om *OrderedMap) Get(key string) (interface{}, bool) {
	value, exists := om.values[key]
	return value, exists
}

// Keys 获取所有键的切片
func (om *OrderedMap) Keys() []string {
	return om.keys
}

// Values 获取所有值的切片
func (om *OrderedMap) Values() []interface{} {
	vals := []interface{}{}
	for _, key := range om.keys {
		vals = append(vals, om.values[key])
	}
	return vals
}

func main() {
	om := NewOrderedMap()
	om.Set("name", "John")
	om.Set("age", 30)
	om.Set("isStudent", false)
	om.Set("scores", []int{90, 85, 92})

	fmt.Println("OrderedMap Keys:", om.Keys())
	fmt.Println("OrderedMap Values:", om.Values())

	if value, exists := om.Get("name"); exists {
		fmt.Println("name:", value)
	}

	if value, exists := om.Get("nationality"); exists {
		fmt.Println("nationality:", value)
	} else {
		fmt.Println("No such key: nationality")
	}

	// 循环打印所有键和值
	fmt.Println("All keys and values in OrderedMap:")
	for _, key := range om.Keys() {
		value, _ := om.Get(key)
		fmt.Printf("%s: %v\n", key, value)
	}
}
