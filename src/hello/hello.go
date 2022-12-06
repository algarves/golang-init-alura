package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const loop = 2
const timeout = 5

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
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Log Viewer")
	fmt.Println("0 - Exit")

	cmd := readCommand()

	fmt.Println("The memory address of the cmd variable is:", &cmd)
	fmt.Println("The input value of the cmd variable is:", cmd)

	fmt.Println("Using conditional:")

	if cmd == 1 {
		startMonitoring()
	} else if cmd == 2 {
		fmt.Println("Logging ...")
	} else if cmd == 0 {
		fmt.Println("Exit")
		exit()
	} else {
		fmt.Println("Command not found.")
		os.Exit(-1)
	}

	fmt.Println("Using switch:")

	switch cmd {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Logging ...")
	case 0:
		fmt.Println("Exit")
		exit()
	default:
		fmt.Println("Command not found.")
		os.Exit(-1)
	}
}

func readCommand() int {
	var cmd int
	fmt.Scan(&cmd)

	return cmd
}

func exit() {
	os.Exit(0)
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br"}

	fmt.Println("Using FOR ...")
	for i := 0; i < len(sites); i++ {
		healthCheck(sites[i], i)
	}

	fmt.Println("")

	fmt.Println("Using FOREACH ...")
	for i := 0; i < loop; i++ {
		for i, site := range sites {
			healthCheck(site, i)
		}
		time.Sleep(timeout * time.Second)
		fmt.Println("")
	}
}

func healthCheck(site string, i int) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Endpoint", i, site, "is UP")
	} else {
		fmt.Println("Endpoint", i, site, "is DOWN")
	}
}
