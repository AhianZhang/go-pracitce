# Golang 修仙之路
语言只是一种工具，其诞生必定是为了解决一类问题，不能拿着锤子看什么都是钉子。程序员最重要的是内功的修炼，痴迷于奇淫技巧往往不会长久。此仓库只是记录学习者的历程，不适用大多读者，望慎之。
## 炼精化气

### 基本类型

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

```

### 变量声明

```
var i, j int = 1, 2
```
函数内类型推断
const 常量不支持类型推断

```
 k := 3
```
### 多返回值

```
func swap(x, y string) (string, string) {
	return y, x
}

```
Naked return statements：
官方不推荐在较长函数中使用，影响可读性
```
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

```

### 类型转换

```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

## 炼气化神
## 炼神还虚
## 炼虚合道