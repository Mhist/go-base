package main

import "fmt"
func main()  {
	months := [...]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "Mary",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December"}
	q1 := months[6:9]
	q2 := months[:3] // 到3不包括3
	q3 := months[3:] // 从3开始包括3
	fmt.Printf("%v\n",months)
	fmt.Printf("%v\n",q1)
	fmt.Printf("%v\n",q2)
	fmt.Printf("%v\n",q3)


	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	fmt.Println(s) // "[1 0 2 3 4 5]"
	reverse(s[2:])
	fmt.Println(s) // "[1 0 5 4 3 2]"
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"


	//内置的len()和cap()函数分别返回slice的长度和容量。
	var summer = []int{}
	if summer != nil { /* ... */
		fmt.Println(false,len(summer),cap(summer)) // "false 0 0"
	}

	var summer2 = []int(nil)
	if summer2 == nil { /* ... */
		fmt.Println(true ,len(summer2),cap(summer2)) // ""
	}

	// append函数
	//内置的append函数用于向slice追加元素：
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"


}
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


