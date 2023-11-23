package main

import (
	"time"
	"fmt"
)

func main() {
	// 时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	now := time.Now()
	fmt.Println(now.In(loc))

	// 使用 Clock() 方法获取时、分、秒
	hour, min, sec := now.Clock()
	fmt.Printf("时：%d 分：%d 秒：%d\n", hour, min, sec)
	// 使用 Hour()、Minute()、Second() 方法获取时、分、秒
	fmt.Printf("时：%d 分：%d 秒：%d\n", now.Hour(), now.Minute(), now.Second())
	year, month, day := now.Date()
	fmt.Printf("年：%d 月：%s 日：%d\n", year, month, day)
	// 使用 Year()、Month()、Day() 方法获取年、月、日
	fmt.Printf("年：%d 月：%s 日：%d\n", now.Year(), now.Month(), now.Day())

	// 在当前时间上加上一小时
	oneHourLater := now.Add(time.Hour)
	fmt.Println("一小时后的时间：", oneHourLater)
	// 在当前时间上加上指定的年、月、日
	oneYearThreeMonthsFiveDaysLater := now.AddDate(1, 3, 5)
	fmt.Println("一年三个月五天后的时间：", oneYearThreeMonthsFiveDaysLater)
	// 计算两个时间的差值
	otherTime := now.Add(2 * time.Hour)
	duration := otherTime.Sub(now)
	fmt.Println("两个时间的差值：", duration)
	// 判断两个时间是否相同
	areEqual := now.Equal(otherTime)
	fmt.Println("两个时间是否相同：", areEqual)
	// 判断时间先后顺序
	isBefore := now.Before(otherTime)
	isAfter := now.After(otherTime)
	fmt.Println("当前时间是否在其他时间之前：", isBefore)
	fmt.Println("当前时间是否在其他时间之后：", isAfter)
	// 计算从某个时间点到现在的时间差
	timeSince := time.Since(now)
	fmt.Println("从当前时间到现在的时间差：", timeSince)
	// 计算从现在到某个时间点的时间差
	timeUntil := time.Until(otherTime)
	fmt.Println("从现在到其他时间的时间差：", timeUntil)
	// 截断时间到指定精度
	truncatedTime := now.Truncate(time.Hour)
	fmt.Println("截断到小时的时间：", truncatedTime)

	// 格式化时间输出
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("格式化时间：", formattedTime)
	// 解析字符串为时间
	timeStr := "2023-11-15 12:30:45"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("解析时间出错：", err)
	} else {
		fmt.Println("解析后的时间：", parsedTime)
	}
	// 解析时间段字符串
	durationStr := "1h30m45s"
	parsedDuration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Println("解析持续时间出错：", err)
	} else {
		fmt.Printf("解析后的持续时间：%v，总小时数：%f\n", parsedDuration, parsedDuration.Hours())
	}
}