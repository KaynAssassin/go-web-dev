package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct{
	Name string
	Motto string
	Admin bool
}

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	u1 := user{
		Name:"Kayn",
		Motto:"Assassin",
		Admin :false,
	}
	u2 := user{
		Name:"Kayn",
		Motto:"Assassin",
		Admin :false,
	}
	users :=  []user{u1,u2}
	err := tpl.Execute(os.Stdout,users)
	if err != nil{
		log.Fatalln(err)
	}
}