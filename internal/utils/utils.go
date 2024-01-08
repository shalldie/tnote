package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 三元运算
func Ternary[T any](condition bool, item1 T, item2 T) T {
	if condition {
		return item1
	}

	return item2
}

// 获取较大的数
func MathMax(a int, b int) int {
	return Ternary(a > b, a, b)
}

// 获取较小的数
func MathMin(a int, b int) int {
	return Ternary(a < b, a, b)
}

func Log(msg string) {

	filePath := "log.out" // 要操作的文件路径

	// 打开文件，如果不存在则创建新文件
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("无法打开文件", filePath, ": ", err)
	}
	defer file.Close()

	contentToAppend := time.Now().Format("2006-01-02 15:04:05") + ": " + msg + "\n" // 要追加的内容

	_, err = file.WriteString(contentToAppend)
	if err != nil {
		log.Fatal("无法向文件追加内容: ", err)
	}

	fmt.Println("成功追加内容到文件")
}
