[Go 语言教程 | 菜鸟教程](https://www.runoob.com/go/go-tutorial.html)

[Go by Example 中文 (studygolang.com)](http://books.studygolang.com/gobyexample/)

# 1.环境准备

1.到[官网](https://go.dev/dl/)下载对应的系统版本

安装完成后可到终端输入

```shell
go version
```

来查看是否安装完毕

关于换源(初始源GOPROXY=https://proxy.golang.org,direct)
```shell
go env 查看当前的配置 # 注意GOPROXY的值 
go env -w GOPROXY=https://goproxy.cn,direct # 永久换源   
```
推荐代理源 
- https://goproxy.cn（七牛云）
- https://goproxy.io（全球代理）
- https://mirrors.aliyun.com/goproxy/（阿里云）

2.下载IDE [goland](https://www.jetbrains.com/go/download/?section=windows)可以选择其他版本,建议zip版

记得在 Settings → Go → Go Modules
✅ Enable Go Modules integration

✅ 勾选 “Download dependencies automatically”

Proxy：保持默认即可

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

## 关闭通道

```go
// 在这个例子中，我们将使用一个 jobs 通道来传递 main() 中 Go协程任务执行的结束信息到一个工作 Go 协程中。
// 当我们没有多余的任务给这个工作 Go 协程时，我们将 close 这个 jobs 通道。
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// for + 二值接收 -> 直到channel关闭并耗尽
	// 这是工作 Go 协程。使用 j, more := <- jobs 循环的从jobs 接收数据。
	// 在接收的这个特殊的二值形式的值中，如果 jobs 已经关闭了，并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。
	// 当我们完成所有的任务时，将使用这个特性通过 done 通道去进行通知。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				// 防止死循环读已关闭的通道
				return
			}
		}
	}()

	//这里使用 jobs 发送 3 个任务到工作函数中，然后关闭 jobs。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 通道同步方法等待任务结束。
	<-done
}
```

## 通道遍历

```go
func main() {

	// 我们将遍历在 queue 通道中的两个值。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 这个 range 迭代从 queue 中得到的每个值。因为我们在前面 close 了这个通道，这个迭代会在接收完 2 个值之后结束。
	//如果我们没有 close 它，我们将在这个循环中继续阻塞执行，等待接收第三个值
	for elem := range queue {
		fmt.Println(elem)
	}
}
```

# 23.定时器

```go
// 我们常常需要在后面一个时刻运行 Go 代码，或者在某段时间间隔内重复运行。Go 的内置 定时器 和 打点器 特性让这些很容易实现。
func main() {

	// 基础定时器(阻塞等待)
	// 创建一个定时器timer1,两秒后触发,内部有一个channel
	timer1 := time.NewTimer(time.Second * 2)
	// <-timer.C 阻塞当前的goroutine(main) 直到两秒后定时器到期 定时器到期时runtime向timer.C发送一个time.time
	// main goroutine被唤醒 相当于time.Sleep(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 可取消的定时器
	// 这段代码实际执行顺序 创建timer2 启动goroutine但是没等到1s 立刻调用timer2.Stop() stop2等于true
	// 定时器被停止,timer2.C永远不会再收到值,goroutine被阻塞 程序结束goroutine被回收
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
```

# 24.打点器

```go
func main() {

	// 打点器和定时器的机制有点相似：用通道来发送数据。这里我们在这个通道上使用内置的 range 来迭代值每隔500ms 发送一次的值。
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// 打点器可以和定时器一样被停止。一旦一个打点停止了，将不能再从它的通道中接收到值。我们将它在运行 1600ms 后停止这个打点器。
	time.Sleep(time.Millisecond * 10000)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

```

# 25.工作池

```go
// 这是我们将要在多个并发实例中支持的任务了。
// 这些执行者将从 jobs 通道接收任务，并且通过 results 发送对应的结果。我们将让每个任务间隔 1s 来模仿一个耗时的任务。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
func main() {

	// 为了使用 worker 工作池并且收集他们的结果，我们需要2 个通道。
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 这里我们发送 9 个 jobs，然后 close 这些通道来表示这些就是所有的任务了。
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后，我们收集所有这些任务的返回值。
	for a := 1; a <= 9; a++ {
		<-results
	}
}

```

# 26.速率控制
```go
func main() {

	// 首先我们将看一下基本的速率限制。假设我们想限制我们接收请求的处理，我们将这些请求发送给一个相同的通道。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 这个 limiter 通道将每 200ms 接收一个值。这个是速率限制任务中的管理器。
	limiter := time.Tick(time.Millisecond * 200)

	// 通过在每次请求前阻塞 limiter 通道的一个接收，我们限制自己每 200ms 执行一次请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们想临时进行速率限制，并且不影响整体的速率控制我们可以通过通道缓冲来实现。这个 burstyLimiter 通道用来进行 3 次临时的脉冲型速率限制。
	burstyLimiter := make(chan time.Time, 3)

	// 想将通道填充需要临时改变3次的值，做好准备。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每 200 ms 我们将添加一个新的值到 burstyLimiter中，直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 现在模拟超过 5 个的接入请求。它们中刚开始的 3 个将由于受 burstyLimiter 的“脉冲”影响。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

```

# 27.原子计数器

```go
func main() {

	// 我们将使用一个无符号整型数来表示（永远是正整数）这个计数器。
	var ops uint64 = 0

	// 为了模拟并发更新，我们启动 50 个 Go 协程，对计数器每隔 1ms 进行一次加一操作。
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 使用 AddUint64 来让计数器自动增加，使用& 语法来给出 ops 的内存地址。
				atomic.AddUint64(&ops, 1)

				// 允许其它 Go 协程的执行
				runtime.Gosched()
			}
		}()
	}

	// 等待一秒，让 ops 的自加操作执行一会。
	time.Sleep(time.Second)

	// 为了在计数器还在被其它 Go 协程更新时，安全的使用它，我们通过 LoadUint64 将当前值的拷贝提取到 opsFinal中。和上面一样，我们需要给这个函数所取值的内存地址 &ops
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
```

# 28.互斥锁

```go
func main() {

	// 在我们的例子中，state 是一个 map。
	var state = make(map[int]int)

	// 这里的 mutex 将同步对 state 的访问。
	var mutex = &sync.Mutex{}

	// 为了比较基于互斥锁的处理方式和我们后面将要看到的其他方式，ops 将记录我们对 state 的操作次数。
	var ops int64 = 0

	// 这里我们运行 100 个 Go 协程来重复读取 state。
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// 每次循环读取，我们使用一个键来进行访问，Lock() 这个 mutex 来确保对 state 的独占访问，读取选定的键的值，Unlock() 这个mutex，并且 ops 值加 1。
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// 为了确保这个 Go 协程不会在调度中饿死，我们在每次操作后明确的使用 runtime.Gosched()进行释放。
				// 这个释放一般是自动处理的，像例如每个通道操作后或者 time.Sleep 的阻塞调用后相似，但是在这个例子中我们需要手动的处理。
				runtime.Gosched()
			}
		}()
	}

	// 同样的，我们运行 10 个 Go 协程来模拟写入操作，使用和读取相同的模式。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	// 让这 10 个 Go 协程对 state 和 mutex 的操作运行 1 s。
	time.Sleep(time.Second)

	// 获取并输出最终的操作计数。
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	// 对 state 使用一个最终的锁，显示它是如何结束的。
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
```

# Go状态协程

```go
func main() {

	// 和前面一样，我们将计算我们执行操作的次数。
	var ops int64

	// reads 和 writes 通道分别将被其他 Go 协程用来发布读和写请求。
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 这个就是拥有 state 的那个 Go 协程，和前面例子中的map一样，不过这里是被这个状态协程私有的。
	// 这个 Go 协程反复响应到达的请求。先响应到达的请求，然后返回一个值到响应通道 resp 来表示操作成功（或者是 reads 中请求的值）
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动 100 个 Go 协程通过 reads 通道发起对 state 所有者Go 协程的读取请求。
	// 每个读取请求需要构造一个 readOp，发送它到 reads 通道中，并通过给定的 resp 通道接收结果。
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 用相同的方法启动 10 个写操作。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 让 Go 协程们跑 1s。
	time.Sleep(time.Second)

	// 最后，获取并报告 ops 值。
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}
```
# 排序
```go
func main() {

	// 排序方法是正对内置数据类型的；这里是一个字符串的例子。
	//注意排序是原地更新的，所以他会改变给定的序列并且不返回一个新值。
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 一个 int 排序的例子。
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// 我们也可以使用 sort 来检查一个序列是不是已经是排好序的。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

```

# 自定义排序
```go
// 为了在 Go 中使用自定义函数进行排序，我们需要一个对应的类型。这里我们创建一个为内置 []string 类型的别名的ByLength 类型，
type ByLength []string

// 我们在类型中实现了 sort.Interface 的 Len，Less和 Swap 方法，这样我们就可以使用 sort 包的通用Sort 方法了，Len 和 Swap 通常在各个类型中都差不多，
// Less 将控制实际的自定义排序逻辑。在我们的例子中，我们想按字符串长度增加的顺序来排序，所以这里使用了 len(s[i]) 和 len(s[j])。
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// 一切都准备好了，我们现在可以通过将原始的 fruits 切片转型成 ByLength 来实现我们的自定排序了。
// 然后对这个转型的切片使用 sort.Sort 方法。
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
```

 # Panic
```go
func main() {
	// panic 意味着有些出乎意料的错误发生。通常我们用它来表示程序正常运行中不应该出现的，或者我们没有处理好的错误。
	// 我们将在这个网站中使用 panic 来检查预期外的错误。这个是唯一一个为 panic 准备的例子。
	panic("a problem")
	// panic 的一个基本用法就是在一个函数返回了错误值但是我们并不知道（或者不想）处理时终止运行。
	//这里是一个在创建一个新文件时返回异常错误时的panic 用法。
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
```
# Defer
Defer 被用来确保一个函数调用在程序执行结束前执行。同样用来执行一些清理工作。 defer 用在像其他语言中的ensure 和 finally用到的地方。
```go
// 假设我们想要创建一个文件，向它进行写操作，然后在结束时关闭它。这里展示了如何通过 defer 来做到这一切。
func main() {
	// 在 closeFile 后得到一个文件对象，我们使用 defer通过 closeFile 来关闭这个文件。
	//这会在封闭函数（main）结束时执行，就是 writeFile 结束后。
	f := createFile("D:\\defer.txt")
	defer closeFile(f)
	writeFile(f)
}
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
```

