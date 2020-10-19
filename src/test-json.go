package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
	"fmt"
)

type Data struct {
	Time time.Time
	Detail string
}

func main(){
	var data1 = Data{
		time.Now(),"",
	}
	fmt.Print("data input =>")
	fmt.Scan(&data1.Detail)
	fmt.Print("\n\n")

	b,err := json.Marshal(data1)
	if err != nil{
		log.Fatal(err)
	}

	file,err:=os.OpenFile("./test.json",os.O_APPEND|os.O_CREATE,0600)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data1)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(b))
	//fmt.Println(data1.Time)
	//fmt.Println(data1.Detail)
}