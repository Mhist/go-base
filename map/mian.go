package main

import (
	"fmt"
	"sort"
)

func main() {
	// 通过内置的函数make创建
	ages := make(map[string]int) // mapping from strings to ints
	fmt.Printf("%v\n", ages)
	// 字面量方式创建
	other := map[string]int{
		"age":  18,
		"year": 2022,
	}
	fmt.Printf("%v\n", other)         // map[age:18 year:2022]
	fmt.Printf("%v\n", other["age"])  // 18
	fmt.Printf("%v\n", other["year"]) // 2022
	// 字面量方式创建空map
	emptyMap := map[string]int{}
	fmt.Printf("%v\n", emptyMap) // map[]
	fmt.Println("a")
	deleteMap := map[string]string{
		"save":    "保留的内容",
		"deleted": "要删除的内容",
		"order_one": "乱序内容1",
		"order_two": "乱序内容2",
		"order_three": "乱序内容3",
	}
	var names [] string
	for name := range deleteMap {
		names = append(names,name)
	}
	fmt.Printf("names排序前：%v\n", names)
	sort.Strings(names)
	fmt.Printf("names排序后：%v\n", names)
	//fmt.Printf("删除前：%v\n", deleteMap) //删除前：map[deleted:要删除的内容 save:保留的内容]
	////delete(deleteMap,"deleted") //
	//fmt.Printf("删除后：%v\n", deleteMap) //删除后：map[save:保留的内容]
	for _, name := range names {
		fmt.Printf("%s\t%s\n", name, deleteMap[name])
	}

	// 零值
	var ageszero map[string]int
	fmt.Println(ageszero == nil)    // "true"
	fmt.Println(len(ageszero) == 0) // "true"

	var agess map[string]string // 空map对于返回key为0，
	agee, ok := agess["bob"]
	if !ok { /* "bob" is not a key in this map; age == 0. */
		fmt.Printf("%T\n",agee)
		fmt.Printf("%v\n",agee)
	}

	var ageeess map[string]int // 空map对于返回值类型得零值，
	ageee, ok := ageeess["bob"]
	if !ok { /* "bob" is not a key in this map; age == 0. */
		fmt.Printf("%T\n",ageee)
		fmt.Printf("%v\n",ageee)
	}

	var ageeess2 map[string]int // 空map对于返回值类型得零值，
	ageee2, ok := ageeess2["bob"]
	if !ok { /* "bob" is not a key in this map; age == 0. */
		fmt.Printf("%T\n",ageee2)
		fmt.Printf("%v\n",ageee2)
	}

	a := equal(ageeess, ageeess2)
	fmt.Printf("比较函数应用%v\n",a)
	//
	//
	//seen := make(map[string]bool) // a set of strings
	//input := bufio.NewScanner(os.Stdin)
	//for input.Scan() {
	//	line := input.Text()
	//	if !seen[line] {
	//		seen[line] = true
	//		fmt.Println(line)
	//	}
	//}
	//
	//if err := input.Err(); err != nil {
	//	fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
	//	os.Exit(1)
	//}

	var ma = make(map[string]int)
	ma["title"] = 15
	ma["age"] = 20
	list := []string{}
	k(list)
	Add(list)
	Count(list)
	fmt.Printf("list%v\n",ma)









}
// 要判断两个map是否包含相同的key和value，我们必须通过一个循环实现：
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
