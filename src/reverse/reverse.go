package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read()(detail []string){

	origin, err := os.Open("./test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer origin.Close()

	string := bufio.NewScanner(origin)

	for i :=1;string.Scan();i++{
		fmt.Println(string.Text())
		detail[i] = string.Text()
	}
	return
}

func write(detail []string){
	origin, err := os.Open("./test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer origin.Close()

	for i:=len(detail);i == 0;i--{
		_,err = origin.Write([]byte(detail[i]))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main(){

	//write(read())

	//var detail []string
	origin, err := os.Open("./test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer origin.Close()

	string := bufio.NewScanner(origin)

	for i :=1;string.Scan();i++{
		fmt.Println(string.Text())	//デバッグ用
		//detail[i] = string.Text()
	}
}