# GoLearn

-----------------------------------------------------

[![last commit](https://img.shields.io/github/last-commit/Coconut-Aero/GoLearn)](https://github.com/Coconut-Aero/GoLearn/commits/master)
[![Static Badge](https://img.shields.io/badge/Coconut-Aero-blue)](https://github.com/Coconut-Aero)

A Repository for Coconut-Aero to learn Go/Golang

_Updated 2024-4-16 16:50 CST_

utils/util-func.go中的各种可以调用的函数

| 描述           |              用法               | 类型   |
|--------------|:-----------------------------:|------|
| 按要求输出一维数组    | PrintArr(arr []int, size int) | void |
| 将一维数组排序      |        Sort(res []int)        | int  |
| 检验一维数组是否排序成功 |     IsSorted(list []int)      | void |



已有项目：

| 序号 | 文件名                   | 描述                          | 备注 |
|----|-----------------------|-----------------------------|----|
| 1  | hello-world.go        | Hello World                 |    |
| 2  | Basic-IO.go           | 基本输入输出                      |    |
| 3  | time.go               | 显示当前系统的UNIX时间               |    |
| 4  | number-pyramid.go     | 数字金字塔                       |    |
| 5  | estimate-pi.go        | 计算圆周率                       |    |
| 6  | array-test.go         | 对数组的操作                      |    |
| 7  | enhanced-random.go    | 增强的random，可以选择范围，排除一些不想要的数字 |    |
| 8  | integer-sum.go        | 命令行参数计算多个数字的和               |    |
| 9  | srpgame.go            | 剪刀石头布游戏                     |    |
| 10 | hangman.go            | 猜单词游戏                       |    |
| 11 | repeated-in-array.go  | 计算数字的出现次数                   |    |

跨文件函数调用记录

| 序号 | 引用自                 | 引用到                    | 引用函数的名称    |
|----|---------------------|------------------------|------------|
| 1  | utils/util-func.go  | sorted.go              | printarr() |
| 2  | utils/util-func.go  | repeated-in-array.go   | Sort()     |