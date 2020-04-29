package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"log"
)

func main(){
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	str := fmt.Sprint(`Hello`+ name + `ahihi`)

	nf,err := os.Create("index.html")

	if err != nil{
		log.Fatal("error creating file",err)
	}

	defer nf.Close()

	io.Copy(nf,strings.NewReader(str))
}