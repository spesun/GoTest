package main

import (
	"fmt"
	"strconv"
)

type ByteCounter int
//该Write方法中，并不Println中的os.Stdout会进行真正的内容输出
//而是仅仅将[]byte的长度保存在ByteCounter中
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {

	testByteCounter()
	testStringer()

}

func testByteCounter()  {
	//测试基本的方法。ByteCounter实现了Write()方法
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	//测试Fprintf中调用Write方法
	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

}

func testStringer()  {
	//测试struct，实现Stringer接口
	var s IntSet2
	s.id = 100
	fmt.Println(s.String())
	fmt.Println((&s).String())

	//这个没异常，即实现了Stringer接口
	var _ fmt.Stringer = &s

	//这个会有异常。应该和下面String()方法定义中，有*IntSet2有关。
	// 如果将*IntSet2，修改成IntSet2，则上、下两种写法都不会有异常。
	//var _ fmt.Stringer = s
}

type IntSet2 struct { id int}
func (b *IntSet2) String() string {return strconv.Itoa(b.id) + "=" + string(b.id)}
