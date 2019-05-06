package testpack;

import (
	"fmt"
)

func Test()  {
	fmt.Print("call  package success ")

	//int 不能和float直接相乘
	var i int = 1;
	var f float64 = 2;
	fmt.Print(float64(i)*f)
}

