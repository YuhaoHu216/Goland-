[Go 语言教程 | 菜鸟教程](https://www.runoob.com/go/go-tutorial.html)

[Go by Example 中文 (studygolang.com)](http://books.studygolang.com/gobyexample/)

# 1.环境准备

1.到[官网](https://go.dev/dl/)下载对应的系统版本

安装完成后可到终端输入

```shell
go version
```

来查看是否安装完毕

2.下载IDE [goland](https://www.jetbrains.com/go/download/?section=windows)可以选择其他版本,建议zip版

# 2.Go语言结构

Go 语言的基础组成有以下几个部分：

- 包声明
- 引入包
- 函数
- 变量
- 语句 & 表达式
- 注释

例如：

```go
package main

import "fmt"

func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")
}
```

1.  *package main* 定义了包名。必须在源文件的第一行指明这个文件属于哪个包。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

2. *import "fmt"* 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。

3. *func main()* 是程序开始执行的函数。main 函数是每一个可执行程序必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。

4.  /*...*/ 是注释，在程序执行时将被忽略。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾，且**不可以嵌套使用**，多行注释一般用于包的文档描述或注释成块的代码片段。

5. 下一行 *fmt.Println(...)* 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
6. 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
7. 在命令行中

```shell
go run xxx.go  # 执行go文件
```

```shell
go build xxx.go # 为go文件生成二进制文件
```

**注意:**`{`不能单独在一行

#  3.Go基础语法

## Go 标记

Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。如以下 GO 语句由 6 个标记组成：

```
fmt.Println("Hello, World!")
```

6 个标记是(每行一个)：

```go
1. fmt
2. .
3. Println
4. (
5. "Hello, World!"
6. )
```

## 行分隔符

在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号`;`结尾,如果把多行代码写在一行就要用`;`隔开,但是不要这么做.

## 注释

注释不会被编译，每一个包应该有相关注释。

单行注释以 `//` 开头的单行注释。多行注释也叫块注释，均已以 `/*` 开头，并以 `*/` 结尾。如：

```go
// 单行注释
/*
 多行注释
 */
```

## 标识符

标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

## 字符串连接

Go 语言的字符串连接可以通过 **+** 实现：

```go
package main
import "fmt"
func main() {
  fmt.Println("Google" + "Runoob")
}
```

## 关键字

下面列举了 Go 代码中会使用到的 25 个关键字或保留字：

| break    | default     | func   | interface | select |
| -------- | ----------- | ------ | --------- | ------ |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：

| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |

程序一般由关键字、常量、变量、运算符、类型和函数组成。

程序中可能会使用到这些分隔符：括号 `()`，中括号 `[]` 和大括号` {}`。

程序中可能会使用到这些标点符号：`.`、`,`、`;`、`:`和 `…`。

## 格式化字符串

Go 语言中使用 **fmt.Sprintf** 或 **fmt.Printf** 格式化字符串并赋值给新串：

- **Sprintf** 根据格式化参数生成格式化的字符串并返回该字符串。
- **Printf** 根据格式化参数生成格式化的字符串并写入标准输出。

```go
package main

import (
  "fmt"
)

func main() {
  // %d 表示整型数字，%s 表示字符串
  var stockcode=123
  var enddate="2025-12-07"
  var url="Code=%d&endDate=%s"
  var target_url=fmt.Sprintf(url,stockcode,enddate)
  fmt.Println(target_url)
  fmt.Printf(url,stockcode,enddate)
}
```

# 4.Go数据类型

- **布尔型**

  布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。

- **数字类型**

  整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。

  - 整型

    - `uint8 uint16 uint32 uint64`

      无符号8/16/32/64整型(0 ~ 2^x)

    - `int8 int16 int32 int64`

      有符号8/16/32/64整型(-2^x ~ 2^x)

  - 浮点型

    - `float32 float64`

      IEEE-754 32/64位浮点型数

    - `complex64 complex128`

      32/64位实数和虚数

  - 其他数字类型

    - `byte` 	    类似uint8
    - `rune`        类似int32
    - `uint`         32或64位
    - `int  `         与uint一样大小
    - `uintptr`   无符号整型,用于存放一个指针

- **字符串类型**

  字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。

- **派生类型**

  - (a) 指针类型（Pointer）
  - (b) 数组类型
  - (c) 结构化类型(struct)
  - (d) Channel 类型
  - (e) 函数类型
  - (f) 切片类型
  - (g) 接口类型（interface）
  - (h) Map 类型

# 5.Go语言变量

声明变量一般形式使用`var`关键字

```go
var identifier type
var identifier1, identifier2 type
```

例如

```go
func main() {
	var a string = "go go go"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
}

```

## 变量声明

**第一种，指定变量类型，如果没有初始化，则变量默认为零值**。

```
var v_name v_type
v_name = value
```

- 数值类型（包括complex64/128）为 **0**

- 布尔类型为 **false**

- 字符串为 **""**（空字符串）

- 以下几种类型为 **nil**：

  ```go
  var a *int
  var a []int
  var a map[string] int
  var a chan int
  var a func(string) int
  var a error // error 是接口
  ```

**第二种，根据值自行判定变量类型。**

```go
var v_name = value
```

**第三种，如果变量已经使用 var 声明过了，再使用 \**:=\** 声明变量，就产生编译错误，格式：**

```
v_name := value
```

例如：

```
var intVal int 
intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
```

直接使用下面的语句即可：

```
intVal := 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
```

**intVal := 1** 相等于：

```go
var intVal int 
intVal = 1 
```

## 多变量声明

```go
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
```

## 值类型和引用类型

所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值

当使用等号 `=` 将一个变量的值赋值给另一个变量时，如：`j = i`，实际上是在内存中将 i 的值进行了拷贝

可以通过 &i 来获取变量 i 的内存地址，值类型变量通常存储在栈中，尤其是当它们是局部变量时。当值类型变量的值需要在函数作用域之外使用时，Go 会将其分配到堆内存中。

更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。

一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置。

这个内存地址称之为指针，这个指针实际上也被存在另外的某一个值中。

同一个引用类型的指针指向的多个字可以是在连续的内存地址中（内存布局是连续的），这也是计算效率最高的一种存储形式；也可以将这些字分散存放在内存中，每个字都指示了下一个字所在的内存地址。

当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。

如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响。

## 注意

如果你在定义变量 a 之前使用它，则会得到编译错误 undefined: a。

如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误 a declared but not used

但是全局变量是允许声明但不使用的。 同一类型的多个变量可以声明在同一行，如：

```go
var a, b, c int
```

多变量可以在同一行进行赋值，如：

```go
var a, b int
var c string
a, b, c = 5, 7, "abc"
```

上面这行假设了变量 a，b 和 c 都已经被声明，否则的话应该这样使用：

```go
a, b, c := 5, 7, "abc"
```

右边的这些值以相同的顺序赋值给左边的变量，所以 a 的值是 5， b 的值是 7，c 的值是 "abc"。

这被称为 并行 或 同时 赋值。

如果你想要交换两个变量的值，则可以简单地使用 **a, b = b, a**，两个变量的类型必须是相同。

空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。

_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。

并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。

# 6.Go常量

量是一个简单值的标识符，在程序运行时，不会被修改的量。

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

用`const`关键字定义常量

```go
const identifier [type] = value
```

你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。

多个相同类型的声明可以简写为：

```go
const c_name1, c_name2 = value1, value2
```

常量还可以用作枚举：

```go
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```

常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：

```go
package main

import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main(){
    println(a, b, c)
}
```

## iota

iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

iota 可以被用作枚举值：

```go
const (
    a = iota
    b = iota
    c = iota
)

```

第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：

```go
const (
    a = iota
    b
    c
)

```

一个用法

```go
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}
```

一个实例

```go
package main

import "fmt"
const (
    i=1<<iota
    j=3<<iota
    k
    l
)

func main() {
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}
```

# 7.Go运算符

算术运算符有(和主流编程语言一致)

`+`  ,`-`  ,`*`  ,`/`  ,`%`  ,`++`  ,`--`

关系运算符(和主流编程语言一致)

`==` ,  `!=`, ` > ` ,  `<`  ,`>=`,  `<=` 

逻辑运算符(和主流编程语言一致)

`&&` , `||`,  `!`

位运算符

`与:&` , `或:|`, `异或:^`, `左移:<<`,  `右移:>>`

赋值运算符

`=`, `+=` , `-=`, `*=`, `/=`, `%=`, `<<=` , `>>=`, `&=`, `^=`, `|=`

其他运算符

`返回变量存储地址:&`, `指针变量:*`

**运算符优先级**

有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

| 优先级 | 运算符           |
| :----- | :--------------- |
| 5      | * / % << >> & &^ |
| 4      | + - \| ^         |
| 3      | == != < <= > >=  |
| 2      | &&               |
| 1      | \|\|             |

可使用小括号来指定优先级

# 8.For循环

`for` 是 Go 中唯一的循环结构。这里有 `for` 循环的三个基本使用方式。

```go
func main(){
	i := 1;
	// 类似于其他语言的while 最常用的方式，带单个循环条件。
	for i <= 10{
		fmt.Println(i)
		i = i + 1
	}
	// 经典的初始化/条件/后续形式 for 循环
	for j := 10; j > 0; j--{
		fmt.Println(j)
	}
	// 不带条件的 for 循环将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环。
	for {
		fmt.Println("loop")
		break;
	}
}
```

# 9.if/else分支

可以不要 `else` 只用 `if` 语句,在条件语句之前可以有一个语句；任何在这里声明的变量都可以在所有的条件分支中使用。在 Go 中，你可以不使用圆括号，但是花括号是需要的。Go 里没有[三目运算符](http://en.wikipedia.org/wiki/%3F:)，所以即使你只需要基本的条件判断，你仍需要使用完整的 `if` 语句.

```go
func main() {
	i := 8
	if i%2 == 0 {
		fmt.Println("i is even")
	} else {
		fmt.Println("i is odd")
	}
	// 条件语句之前可以有一个语句；任何在这里声明的变量都可以在所有的条件分支中使用
	if num := 9; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}
}
```

# 10.switch 分支结构

```go
func main() {
	i := 2
	fmt.Println("write i as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}
	// default是所有case都不匹配的操作
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式。这里展示了 case 表达式是如何使用非常量的
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	case t.Hour() >= 12:
		fmt.Println("it's after noon")
	}
}
```

# 11.数组

```go
func main() {
	var a [5]int
	// 不初始化,默认零值,int的零值就是0
	fmt.Println("emp:", a)
	// 可将数组中的值拿出来单独操作
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	// 使用内置len函数获取数组的长度
	fmt.Println("len:", len(a))
	// 创建时可初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j

		}
	}
	fmt.Println("2d: ", twoD)
}
```

# 12.slice

```go
func main() {
	// 使用内建的方法 make。这里我们创建了一个长度为3的 string 类型 slice（初始化为零值）
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// 操纵和数组类似
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	// slice可复制,但是必须相同长度
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 这里包含s[2],s[3],s[4]三个元素
	l := s[2:5]
	fmt.Println("sl1:", l)
	// 从s[0]包含到s[5]
	l = s[:5]
	fmt.Println("sl2:", l)
	// 从s[2]到最后一个值
	l = s[2:]
	fmt.Println("sl3:", l)
	// 初始化时赋值
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// 和组成多维数组结构,内部slice长度可以不同,和多维数组不同
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
```

# 13.关联数组(map)

```go
func main() {
	//  map，需要使用内建函数 make(map[key-type]val-type).
	m := make(map[string]int)
	// 使用典型的 make[key] = val 语法来设置键值对。

	m["k1"] = 7
	m["k2"] = 13
	// 使用例如 Println 来打印一个 map 将会输出所有的键值对。
	fmt.Println("map:", m)

	// 使用 name[key] 来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// 当对一个 map 调用内建的 len 时，返回的是键值对数目
	fmt.Println("len:", len(m))

	// 内建的 delete 可以从一个 map 中移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 你也可以通过这个语法在同一行申明和初始化一个新的map。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
```

# 14.range便利(增强for)

```go
func main() {
	// 这里我们使用 range 来统计一个 slice 的元素个数。数组也可以采用这种方法。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range 在数组和 slice 中都同样提供每个项的索引和值。上面我们不需要索引，所以我们使用 空值定义符_ 来忽略它。有时候我们实际上是需要这个索引的。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range 在 map 中迭代键值对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range 在字符串中迭代 unicode 编码。第一个返回值是rune 的起始字节位置，然后第二个是 rune 自己。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
```

# 15.函数

```go
// 这里是一个函数，接受两个 int 并且以 int 返回它们的和
func plus(a int, b int) int {
	// Go 需要明确的返回值，例如，它不会自动返回最后一个表达式的值
	return a + b
}

// (int, int) 在这个函数中标志着这个函数返回 2 个 int。
func vals() (int, int) {
	return 3, 7
}

// 这个函数使用任意数目的 int 作为参数。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 这个 intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数。这个返回的函数使用闭包的方式 隐藏 变量 i。
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// face 函数在到达 face(0) 前一直调用自身。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	// 通过funcName(args) 来调用一个函数，
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	// 这里我们通过多赋值 操作来使用这两个不同的返回值。
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	// 如果你仅仅想返回值的一部分的话，你可以使用空白定义符 _。
	_, c := vals()
	fmt.Println(c)

	// 变参函数使用常规的调用方式，除了参数比较特殊。
	sum(1, 2)
	sum(1, 2, 3)
	// 如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)。
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// 我们调用 intSeq 函数，将返回值（也是一个函数）赋给nextInt。这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时都会更新 i 的值。
	nextInt := intSeq()
	//通过多次调用 nextInt 来看看闭包的效果。
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// 为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下。
	newInts := intSeq()
	fmt.Println(newInts())

	// 看看递归函数的结果
	fmt.Println(fact(7))
}
```

# 16.指针

```go
// 通过两个函数：zeroval 和 zeroptr 来比较指针和值类型的不同。
// zeroval 有一个 int 型参数，所以使用值传递。zeroval 将从调用它的那个函数中得到一个 ival形参的拷贝。
func zeroval(ival int) {
	ival = 0
}

// zeroptr 有和上面不同的 *int 参数，意味着它用了一个 int指针。
// 函数体内的 *iptr 接着解引用 这个指针，从它内存地址得到这个地址对应的当前值。
// 对一个解引用的指针赋值将会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
	*iptr = 0
}
func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	// 通过 &i 语法来取得 i 的内存地址，例如一个变量i 的指针。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针也是可以被打印的。
	fmt.Println("pointer:", &i)

	// zeroval 在 main 函数中不能改变 i 的值，但是zeroptr 可以，因为它有一个这个变量的内存地址的引用。
}

```

# 17.结构体

```go
// 这里的 person 结构体包含了 name 和 age 两个字段。
type person struct {
	name string
	age  int
}

func main() {
	// 使用这个语法创建了一个新的结构体元素。
	fmt.Println(person{"Bob", 20})

	// 你可以在初始化一个结构体元素时指定字段名字。
	fmt.Println(person{name: "Alice", age: 30})

	// 省略的字段将被初始化为零值。
	fmt.Println(person{name: "Fred"})

	// & 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40})

	// 使用点来访问结构体字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 也可以对结构体指针使用. - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)

	// 结构体是可变的。
	sp.age = 51
	fmt.Println(sp.age)
}
```

# 18.方法

```go
type rect struct {
	width, height int
}

// 这里的 area 方法有一个接收器类型 rect。
func (r *rect) area() int {
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收器定义方法。这里是一个值类型接收器的例子。
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
func main() {
	r := rect{width: 10, height: 5}

	// 这里我们调用上面为结构体定义的两个方法。
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go 自动处理方法调用时的值和指针之间的转化
	// 可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
} 
```

# 19.接口

```go
// 这里是一个几何体的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

// 让 rect2 和 circle 实现这个接口
type rect2 struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 在Go中实现一个接口，我们只需要实现接口中的所有方法。这里我们让 rect2 实现了 geometry 接口。
func (r rect2) area() float64 {
	return r.width * r.height
}
func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

// circle 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量的是接口类型，那么我们可以调用这个被命名的接口中的方法。这里有一个一通用的 measure 函数，利用这个特性，它可以用在任何 geometry 上。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	r := rect2{width: 3, height: 4}
	c := circle{radius: 5}
	// 结构体类型 circle 和 rect2 都实现了 geometry接口，所以我们可以使用它们的实例作为 measure 的参数。
	measure(r)
	measure(c)
}
```

# 20.错误处理

```go
// 按照惯例，错误通常是最后一个返回值并且是 error 类型，一个内建的接口。
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New 构造一个使用给定的错误信息的基本error 值。
		return -1, errors.New("can't work with 42")
	}
	// 返回错误值为 nil 代表没有错误。
	return arg + 3, nil
}

// 通过实现 Error 方法来自定义 error 类型是可以的。这里使用自定义错误类型来表示上面的参数错误。
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int) (int, error) {
	if arg == 42 {
		// 在这个例子中，我们使用 &argError 语法来建立一个新的结构体，并提供了 arg 和 prob 这个两个字段的值。
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
func main() {
	// 下面的两个循环测试了各个返回错误的函数。注意在if行内的错误检查代码，在Go中是一个普遍的用法。
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 你如果想在程序中使用一个自定义错误类型中的数据，你需要通过类型断言来得到这个错误类型的实例。
	// 类型断言:如果错误类型属于argError就用ae接收转换后的*argError,否则就是nil,只有ok=true时才进入代码块
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
```

# 21.协程

```go
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {
	// 假设我们有一个函数叫做 f(s)。我们使用一般的方式调并同时运行。
	f("direct")

	// 使用 go f(s) 在一个 Go 协程中调用这个函数。这个新的 Go 协程将会并行的执行这个函数调用。
	go f("goroutine")

	// 你也可以为匿名函数启动一个 Go 协程。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在这两个 Go 协程在独立的 Go 协程中异步的运行，所以我们需要等它们执行结束。
	// 这里的 Scanln 代码需要我们在程序退出前按下任意键结束。
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
```

# 22.通道

## 通道入门

```go
func main() {
	// 使用 make(chan val-type) 创建一个新的通道。通道类型就是他们需要传递值的类型。
	messages := make(chan string)

	// 使用 channel <- 语法 发送 一个新的值到通道中。
	//这里我们在一个新的 Go 协程中发送 "ping" 到上面创建的messages 通道中。
	go func() { messages <- "ping" }()

	// 使用 <-channel 语法从通道中 接收 一个值。这里将接收我们在上面发送的 "ping" 消息并打印出来。
	msg := <-messages
	fmt.Println(msg)

	/**
	我们运行程序时，通过通道，消息 "ping" 成功的从一个 Go 协程传到另一个中。
	默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。
	这个特性允许我们，不使用任何其它的同步操作，来在程序结尾等待消息 "ping"。
	*/
}
```

## 通道缓冲

```go
func main() {
	// 这里我们 make 了一个通道，最多允许缓存 2 个值。
	messages := make(chan string, 2)

	// 因为这个通道是有缓冲区的，即使没有一个对应的并发接收方，我们仍然可以发送这些值。
	messages <- "buffered"
	messages <- "channel"

	// 然后我们可以像前面一样接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
```

## 通道同步

```go
// done 通道将被用于通知其他 Go 协程这个函数已经工作完毕。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// 发送一个值来通知我们已经完工啦。
	done <- true
}
func main() {
	// 运行一个 worker Go协程，并给予用于通知的通道。
	done := make(chan bool, 1)
	go worker(done)
	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	<-done
	/**
	把 <- done 这行代码从程序中移除，程序甚至会在 worker还没开始运行时就结束了。
	main 函数所在的 goroutine 一结束，整个程序就立刻退出，所有其他 goroutine 都会被强制终止，不管它们有没有开始执行。
	<-done 的作用，就是阻塞 main goroutine，防止程序提前退出。
	*/
}
```

## 通道方向

```go
// ping 函数定义了一个只允许发送数据的通道。尝试使用这个通道来接收数据将会得到一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong 函数允许通道（pings）来接收数据，另一通道（pongs）来发送数据。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
```

## 通道选择器

```go
func main() {

	// 在我们的例子中，我们将从两个通道中选择。
	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在若干时间后接收一个值，这个用来模拟例如并行的 Go 协程中阻塞的 RPC 操作
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// 我们使用 select 关键字来同时等待这两个值，并打印各自接收到的值。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
```

## 超时处理

```go
/*
超时 对于一个连接外部资源，或者其它一些需要花费执行时间的操作的程序而言是很重要的。
得益于通道和 select，在 Go中实现超时操作是简洁而优雅的。
*/
func main() {

	//在我们的例子中，假如我们执行一个外部调用，并在 2 秒后通过通道 c1 返回它的执行结果。
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	//这里是使用 select 实现一个超时操作。res := <- c1 等待结果，<-Time.After 等待超时时间 1 秒后发送的值。
	//由于 select 默认处理第一个已准备好的接收操作，如果这个操作超过了允许的 1 秒的话，将会执行超时 case。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	// 如果我允许一个长一点的超时时间 3 秒，将会成功的从 c2接收到值，并且打印出结果。
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
```

## 非阻塞通道操作

```go
/*
常规的通过通道发送和接收数据是阻塞的。
然而，我们可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。
*/
// 通俗讲就是select会马上检查所有case,如果有可以执行的就执行一个,都不能执行就走default的逻辑,不会等待
func main() {
	// 都是无缓冲的channel,当前程序中没有任何goroutine向其发送数据,所有接收和发送都会阻塞
	messages := make(chan string)
	signals := make(chan bool)

	// 这里是一个非阻塞 接收 的例子。如果在 messages 中存在值，然后 select 将这个值带入 <-messages case中。
	// 如果不存在，就直接到 default 分支中。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 一个非阻塞 发送 的实现方法和上面一样。
	// 如果没有default就会一直卡(阻塞)在messages <- msg
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。
	// 这里我们试图在 messages和 signals 上同时使用非阻塞的接受操作。
	// 这里如果msg := <-messages和sig := <-signals都可执行的话就会随机选择一个执行
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
```

