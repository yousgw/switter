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
		//fmt.Println(string.Text())
		detail = append(detail, string.Text()+"\n")
		//fmt.Println(detail[i-1])
	}
	return
}

func write(detail []string){
	origin, err := os.OpenFile("./test.json",os.O_WRONLY,0600)
	if err != nil {
		log.Fatal(err)
	}
	defer origin.Close()

	for i:=len(detail)-1;i != -1;i--{
		_,err = origin.Write([]byte(detail[i]))
		fmt.Print(detail[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main(){

	write(read())

}