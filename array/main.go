package main

import (
	"crypto/sha256"
	"fmt"
)
func main()  {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q,r,r[2]) // // "0"

	q2 := [...]int{1, 2, 3}
	fmt.Printf("%v\n", q2) // "[3]int"
	fmt.Printf("%T\n", q2) // "[3]int"

	// 通过new关键字
	var d = new([10]int)
	fmt.Printf("%v\n", d) //

	//
	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(USD, symbol) // "0 [$ € ￡ ￥]"
	fmt.Println(EUR, symbol) // "1 [$ € ￡ ￥]"
	fmt.Println(GBP, symbol) // "2 [$ € ￡ ￥]"
	fmt.Println(RMB, symbol) // "3 [$ € ￡ ￥]"
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	r2 := [...]int{99: -1}
	fmt.Println(r2) // ""

	// 数组的比较
	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1:= [2]int{1, 3}
	fmt.Println(a1 == b1, a1 == c1, b1 == c1) // "true false false"
	d1 := [2]int{1, 2}
	//d2 := [3]int{1, 2}
	fmt.Println(a1) // compile error: cannot compare [2]int == [3]int
	fmt.Println(d1)
	fmt.Println(d1 == a1)
	//fmt.Println(d2 == a1)

//https://www.cnblogs.com/white-album2/p/16018590.html
	// 练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
	h1 := sha256.Sum256([]byte("x"))
	h2 := sha256.Sum256([]byte("Y"))

	fmt.Printf("%x\n%x\n%x\n", h1,len(h1),h2)
	fmt.Println(popCount(h1, h2))
}
var pc [256]byte
func popCount(s1, s2 [32]byte) int {

	count := 0
	for i := 0; i < 32; i++ {
		temp := s1[i] ^ s2[i]
		count += int(pc[temp])
	}

	return count
}