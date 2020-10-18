package main

import (
	"time"
	"fmt"
)

type Data struct {
	Time time.Time
	Detail string
}

func main(){
	var data1 = Data{
		time.Now(),"test1",
	}
	fmt.Println(data1.Time)
	fmt.Println(data1.Detail)
}