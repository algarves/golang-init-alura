package main

import (
	"fmt"
	"reflect"
)

func main() {
	hello()

	fmt.Println("---")

	variablesEx1()

	fmt.Println("---")

	variablesEx2()

	fmt.Println("---")

	command()
}

func hello() {
	fmt.Println("Hello World")
}

func variablesEx1() {
	var name string = "Khristian"
	var age int = 38
	var version float32 = 1.1

	fmt.Println("Hello", name, "your age is", age)
	fmt.Println("This program is in version", version)
}

func variablesEx2() {
	name := "Khristian"
	age := 38
	version := 1.1

	fmt.Println("The value of", name, "is", reflect.TypeOf(name))
	fmt.Println("The value of", age, "is", reflect.TypeOf(age))
	fmt.Println("The value of", version, "is", reflect.TypeOf(version))
}

func command() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var cmd int
	fmt.Scan(&cmd)

	fmt.Println("The memory address of the cmd variable is:", &cmd)
	fmt.Println("The input value of the cmd variable is:", cmd)
}
