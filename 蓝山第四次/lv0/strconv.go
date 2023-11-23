package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转换为整数
	intStr := "123"
	intValue, err := strconv.Atoi(intStr)
	if err != nil {
		fmt.Println("字符串转整数出错:", err)
	} else {
		fmt.Println("整数值:", intValue)
	}
	// 字符串转换为整数（指定进制）
	hexStr := "1a"
	hexValue, err := strconv.ParseInt(hexStr, 16, 0)
	if err != nil {
		fmt.Println("16进制字符串转整数出错:", err)
	} else {
		fmt.Println("16进制整数值:", hexValue)
	}
	// 字符串转换为无符号整数（指定进制）
	uintStr := "255"
	uintValue, err := strconv.ParseUint(uintStr, 10, 0)
	if err != nil {
		fmt.Println("字符串转无符号整数出错:", err)
	} else {
		fmt.Println("无符号整数值:", uintValue)
	}
	// 字符串转换为浮点数
	floatStr := "3.14"
	floatValue, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("字符串转浮点数出错:", err)
	} else {
		fmt.Println("浮点数值:", floatValue)
	}
	// 字符串转换为布尔值
	boolStr := "true"
	boolValue, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("字符串转布尔值出错:", err)
	} else {
		fmt.Println("布尔值:", boolValue)
	}

	// 整数转换为字符串
	intValueToStr := 42
	intStrFromInt := strconv.Itoa(intValueToStr)
	fmt.Println("整数转字符串:", intStrFromInt)
	// 格式化整数为字符串
	otherIntValueToStr := 12345
	otherIntStrFromInt := strconv.FormatInt(int64(otherIntValueToStr), 10)
	fmt.Println("格式化整数为字符串:", otherIntStrFromInt)
	// 格式化无符号整数为字符串
	uintValueToStr := uint64(98765)
	uintStrFromUint := strconv.FormatUint(uintValueToStr, 10)
	fmt.Println("格式化无符号整数为字符串:", uintStrFromUint)
	// 格式化浮点数为字符串
	floatValueToStr := 2.71828
	floatStrFromFloat := strconv.FormatFloat(floatValueToStr, 'f', -1, 64)
	fmt.Println("格式化浮点数为字符串:", floatStrFromFloat)
	// 布尔值转换为字符串
	boolValueToStr := true
	boolStrFromBool := strconv.FormatBool(boolValueToStr)
	fmt.Println("布尔值转字符串:", boolStrFromBool)
}
