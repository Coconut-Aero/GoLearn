package main

import "fmt"

func main() {
	var name string
	var birthmonth int
	var birthday int
	fmt.Print("输入名字：")
	fmt.Scan(&name)
	fmt.Print("输入出生月：")
	fmt.Scan(&birthmonth)
	fmt.Print("输入出生日：")
	fmt.Scan(&birthday)
	fmt.Print(name, "出生于", birthmonth, "月", birthday, "日")
}
