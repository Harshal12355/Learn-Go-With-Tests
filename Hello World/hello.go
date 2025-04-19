package main

import (
	"fmt"
	"strings"
  )

func Hello_World() string {
	return "Hello World"
}

func Hello_You(name string) string {
	if name == ""{
		name = "World"
	}
	return "Hello " + name 
}

func greeting(language string) (prefix string){
	lang := strings.ToLower(language)
	switch lang {
		case "spanish":
			prefix = "Hola "
		case "english":
			prefix = "Hello "
		case "french":
			prefix = "Bonjour "
		default  :
			prefix = "Hello "
	}
	return 
}
func Hello_lang(name, language string) string {
	if name == ""{
		name = "World"
	}

	return greeting(language) + name
}

func main(){
	fmt.Println("Hello World")
}