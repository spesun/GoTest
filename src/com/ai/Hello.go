package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
	"com/ai/testpack"

) //引入fmt库

func main() {

	a1 := 1
	b1 := 2
	if a1 != 2 && b1 != 1 {
		fmt.Print("ddd")
	} else {
		fmt.Print("ccc")
	}

	testPackage()

	var b string //等同于  b: = ""
	fmt.Println(b)

	var a = "Hello World!" + "aaa"
	fmt.Println(a)

	// 数组循环
	testArray()

	testMap()

	//
	//scanInput()

	//slice测试
	testSlice();

	//funcation测试
	testFunctionParam()

	//function array
	funArray(1, "2")

	//指针
	testPoint()
	testFuncPoint2()

	//指针
	var pointTest PointTest = 1
	//pointTest  := new(PointTest) //可以使用new, 返回值是地址
	pointTest.add() // 会改变bb的值.调用时不是必须写&，go会自动进行转换
	fmt.Println(pointTest)
	pointTest.add2(); //不会改变bb的值. go中函数参数是值传递
	fmt.Println(pointTest)

}

func testPackage()  {
	testpack.Test()
}

//array
func testArray()  {
	fmt.Println("========= array 测试")
	var arr = []string{"0","3","4"}
	var s, sep string
	for i := 0; i < len(arr) ; i++ {  //for initialization; condition; post
		s += sep + arr[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("=============")

	// string或者slice类型值的循环
	for i, arg := range arr[1:] {  //空白标识符_可以在任何你 需要接收自己不想处理的值时使用。
		fmt.Println(  strconv.Itoa(i)  + "=" + arg)
	}

	fmt.Println(strings.Join(arr, " "))
}

//map
func testMap()  {
	fmt.Println("===== map测试")
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages)

}

//slice
func  testSlice()  {
	fmt.Println("===== slice 测试")
	ss :=[]int{1,2,3 }
	runes := append(ss, 3);
	fmt.Println(runes)
}

//function
func testFunctionParam()  {
	fmt.Println("===== 以function作为参数的测试")
	fmt.Println(strings.Map(add1, "HAL-9000"))
}
func add1(r rune) rune { return r + 1 }

//===========
func funArray(arr ...interface{})  {
	fmt.Println("===== 以interface作为参数的测试")
	for _,a := range arr{
		fmt.Println(a)
	}
}

//系统输入
func testScanInput()  {

	fmt.Println("============= scan input")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		//等同于
		//line := input.Text()
		//counts[line] = counts[line] + 1
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//指针测试
func testPoint()  {
	fmt.Println("====== 指针测试")
	a := 1
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
}

//=======
func testFuncPoint2()  {
	fmt.Println("====== 指针作为参数的测试")
	var intSet IntSet;
	intSet.id = 100
	var ast = intSet.String()
	fmt.Println(ast)

}
type IntSet struct { id int}
func (b *IntSet) String() string {return strconv.Itoa(b.id) + "=" + string(b.id)}

// 函数指针测试
type PointTest int
func (c *PointTest) add()  {
	*c = *c + 1
}
func (c PointTest) add2()  {
	c = c + 1
}

