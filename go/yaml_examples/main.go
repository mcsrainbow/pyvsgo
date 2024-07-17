package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	// 创建记录器，不使用 log.Lshortfile 标志
	InfoLogger = log.New(os.Stdout, "", 0)
	ErrorLogger = log.New(os.Stderr, "", 0)
}

func logWithLevel(logger *log.Logger, level string, msg string, v ...interface{}) {
	// 获取当前时间
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	// 构建日志消息
	prefix := fmt.Sprintf("%s %s: ", timestamp, level)
	logger.SetPrefix(prefix)
	logger.Printf(msg, v...)
}

func main() {
	tmpFile := "tmp/yaml_examples.yaml"

	// 生成 YAML 文件的数据
	data := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]interface{}{
			"subkey1": "subvalue1",
			"subkey2": "subvalue2",
		},
		"key3": []interface{}{1, 2, 3},
	}

	logWithLevel(InfoLogger, "INFO", "生成 YAML 的初始数据:")
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		logWithLevel(ErrorLogger, "ERROR", "JSON 序列化失败: %v", err)
		return
	}
	fmt.Println(string(out))

	// 将数据写入 YAML 文件
	out, err = yaml.Marshal(data)
	if err != nil {
		logWithLevel(ErrorLogger, "ERROR", "YAML 序列化失败: %v", err)
		return
	}
	err = os.WriteFile(tmpFile, out, 0644)
	if err != nil {
		logWithLevel(ErrorLogger, "ERROR", "写入 YAML 文件失败: %v", err)
		return
	}

	// 更新 key2.subkey1 数据
	data["key2"].(map[string]interface{})["subkey1"] = "new_subvalue1"

	// 将更新后的数据写回 YAML 文件
	out, err = yaml.Marshal(data)
	if err != nil {
		logWithLevel(ErrorLogger, "ERROR", "YAML 序列化失败: %v", err)
		return
	}
	err = os.WriteFile(tmpFile, out, 0644)
	if err != nil {
		logWithLevel(ErrorLogger, "ERROR", "更新 YAML 文件失败: %v", err)
		return
	}
	logWithLevel(InfoLogger, "INFO", "更新 YAML key2.subkey1: %v", data["key2"].(map[string]interface{})["subkey1"])

	// 读取 YAML 文件以验证更新
	file, err := os.ReadFile(tmpFile)
	if err != nil {
		logWithLevel(ErrorLogger, "ERROR", "读取 YAML 文件失败: %v", err)
		return
	}

	var updatedData map[string]interface{}
	err = yaml.Unmarshal(file, &updatedData)
	if err != nil {
		logWithLevel(ErrorLogger, "ERROR", "解析 YAML 文件失败: %v", err)
		return
	}

	logWithLevel(InfoLogger, "INFO", "读取更新后的 YAML 文件内容 'tmp/yaml_examples.yaml':")
	out, err = yaml.Marshal(updatedData)
	if err != nil {
		logWithLevel(ErrorLogger, "ERROR", "YAML 序列化失败: %v", err)
		return
	}
	fmt.Println(string(out))

	// 校验更新是否成功
	if updatedData["key2"].(map[string]interface{})["subkey1"] != "new_subvalue1" {
		logWithLevel(ErrorLogger, "ERROR", "ERROR: 更新失败")
		return
	}
	logWithLevel(InfoLogger, "INFO", "更新校验成功")
}

