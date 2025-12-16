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

