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

func write_json(data Data) {

	file, err := os.OpenFile("./test.json", os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

func read_json()(string *bufio.Scanner){

	datas, err := os.Open("./test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer datas.Close()
	string = bufio.NewScanner(datas)
	return
}

func output_data(string *bufio.Scanner){
	var data Data
	var now, iftime time.Time
	var duration  time.Duration
	for i :=1;string.Scan();i++{
		json.Unmarshal([]byte(string.Text()),&data)
		duration = now.Sub(data.Time)

		switch {
		case now.Year() - data.Time.Year() != 0://年数が異なる
			iftime = data.Time
			iftime = iftime.AddDate(now.Year() - data.Time.Year(),0,0)
			if now.Sub(iftime).Milliseconds() >= 0{//1年以上前
				fmt.Println(now.Year() - data.Time.Year(),"年前")
			}else {
				if now.Year()-data.Time.Year()-1 != 0 { //ex)年は3ずれているが2年前
					fmt.Println(now.Year()-data.Time.Year()-1, "年前")
				} else {
					goto nextCase //1年以内
				}
			}
			break
		nextCase:
			fallthrough
		case duration.Hours() >= 24://1日以上前
			fmt.Println(data.Time.Month(),"月",data.Time.Day(),"日")
		}


		//fmt.Println(data.Time.Year(),"年",data.Time.Month(),"月",data.Time.Day(),"日",data.Time.Hour(),":",data.Time.Minute(),":",data.Time.Second())
		fmt.Println(data.Detail,"\n")
	}
}

func main() {
	var data1 = Data{
		time.Now(), "",
	}
	fmt.Print("data input =>")
	fmt.Scan(&data1.Detail)
	fmt.Print("\n\n")

	write_json(data1)

	read_json()


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
	/*
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
	*/

}
