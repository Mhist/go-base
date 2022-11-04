package main

import (
	"fmt"
	"time"
)
func main()  {
	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}

	var dilbert Employee

	dilbert.ID = 202210120001
	dilbert.Name = "JYD"
	dilbert.Address = "湖北省武汉市"
	dilbert.DoB = time.Now()
	dilbert.Position = "中国"
	dilbert.Salary = 12000
	dilbert.ManagerID = 2207024

	BID := &dilbert.ID
	*BID = 12232 + *BID
	fmt.Printf("%v\n",*BID)


}