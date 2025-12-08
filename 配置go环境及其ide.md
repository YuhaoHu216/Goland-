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

# 4.Go 语言数据类型

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