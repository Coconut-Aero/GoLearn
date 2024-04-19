# GoLearn

-----------------------------------------------------

[![last commit](https://img.shields.io/github/last-commit/Coconut-Aero/GoLearn)](https://github.com/Coconut-Aero/GoLearn/commits/master)
[![Static Badge](https://img.shields.io/badge/Coconut-Aero-blue)](https://github.com/Coconut-Aero)

A Repository for Coconut-Aero to learn Go/Golang

[![OrelaV](https://i2.hdslb.com/bfs/face/d812a48f1ca84d4f60a112dc31ba65546a787a76.jpg@240w_240h_1c_1s_!web-avatar-space-header.avif "@OrelaV")](https://space.bilibili.com/3546375738361934)

关注[OrelaV](https://space.bilibili.com/3546375738361934)谢谢喵

_Updated 2024-4-19 21:00 CST_

utils/util-func.go中的各种可以调用的函数

| 描述                           |                      用法                       | 类型   |
|------------------------------|:---------------------------------------------:|------|
| 按要求输出一维数组                    |         PrintArr(arr []int, size int)         | void |
| 将一维数组排序                      |                Sort(res []int)                | int  |
| 检验一维数组是否排序成功                 |             IsSorted(list []int)              | void |
| 增强的random，可以选择范围，排除一些不想要的数字  | GetRandom(start int, end int, numbers ...int) | int  |



已有项目：

| 序号 | 文件名                     | 描述                          | 备注                        |
|----|-------------------------|-----------------------------|---------------------------|
| 1  | hello-world.go          | Hello World                 |                           |
| 2  | Basic-IO.go             | 基本输入输出                      |                           |
| 3  | time.go                 | 显示当前系统的UNIX时间               | 复刻time.java               |
| 4  | number-pyramid.go       | 数字金字塔                       | 复刻numberpyramid.java      |
| 5  | estimate-pi.go          | 计算圆周率                       | 复刻estimatepi.java         |
| 6  | array-test.go           | 对数组的操作                      | 复刻arraytest.java          |
| 7  | enhanced-random.go      | 增强的random，可以选择范围，排除一些不想要的数字 | 复刻enhancedRandom.java     |
| 8  | integer-sum.go          | 命令行参数计算多个数字的和               | 复刻intSum.java             |
| 9  | srpgame.go              | 剪刀石头布游戏                     | 复刻srpgame.java            |
| 10 | hangman.go              | 猜单词游戏                       | 复刻hangman.java            |
| 11 | repeated-in-array.go    | 计算数字的出现次数                   | 复刻repeatedInArray.java    |
| 12 | single-digit-number.go  | 由100个个位数组成的数组中各个数字出现的次数统计   | 复刻numberSingleDigit.java  |

跨文件函数调用记录

| 序号 | 引用自                  | 引用到                     | 引用函数的名称     |
|----|----------------------|-------------------------|-------------|
| 1  | utils/util-func.go   | sorted.go               | PrintArr()  |
| 2  | utils/util-func.go   | repeated-in-array.go    | Sort()      |
| 3  | utils/util-func.go   | single-digit-number.go  | GetRandom() |