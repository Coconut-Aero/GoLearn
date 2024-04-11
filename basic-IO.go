package main

import "fmt"

func main() {
	var name string
	var birthmonth int
	var birthday int
	fmt.Print("输入名字：")
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}
	fmt.Print("输入出生月：")
	_, err2 := fmt.Scan(&birthmonth)
	if err2 != nil {
		return
	}
	fmt.Print("输入出生日：")
	_, err3 := fmt.Scan(&birthday)
	if err3 != nil {
		return
	}
	fmt.Print(name, "出生于", birthmonth, "月", birthday, "日")
}
