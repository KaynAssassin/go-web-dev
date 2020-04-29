package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct{
	Name string
	Motto string
}

type car struct {
	Manufacturer string
	Model string
	Doors int
}

var fm = template.FuncMap{
	"uc":strings.ToUpper,
	"ft" : firstThree,
}

func init(){
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string{
	s = strings.TrimSpace(s)
	if len(s) >= 3{
		s = s[:3]
	}

	return s
}

func main(){
	b := sage{
		Name:"Kayn",
		Motto:"Assassin",
	}

	g := sage{
		Name:"Kayn",
		Motto:"Darkin",
	}


	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b,g,m}
	cars := []car{f,c}

	data := struct{
		Wisdom []sage
		Transport []car
	}{
		sages,
		cars,
	}
	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",data)

	if err != nil{
		log.Fatalln(err)
	}
}