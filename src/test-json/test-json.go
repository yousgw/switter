package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Data struct {
	Time   time.Time
	Detail string
}

func main() {
	var data1 = Data{
		time.Now(), "",
	}
	fmt.Print("data input =>")
	fmt.Scan(&data1.Detail)
	fmt.Print("\n\n")

	_, err := json.Marshal(data1)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(data1)

	//var data3 Data
	//json.Unmarshal(jdata,&data3)

	//書き込みテスト
	/*
	file, err := os.OpenFile("./test.json", os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data1)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	*/

	//読み出しテスト

	var data2 Data

	datas, err := os.Open("./test.json")
	if err != nil {
		log.Fatal(err)
	}
	string := bufio.NewScanner(datas)

	for i :=1;string.Scan();i++{
		//fmt.Println(string.Text())
		json.Unmarshal([]byte(string.Text()),&data2)
		fmt.Println(data2.Time.Year(),"年",data2.Time.Month(),"月",data2.Time.Day(),"日",data2.Time.Hour(),":",data2.Time.Minute(),":",data2.Time.Second())
		fmt.Println(data2.Detail,"\n")
	}


}
