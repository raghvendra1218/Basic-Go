package main

import "fmt"

func main() {
	//Declare and assign a very simple map
	color := map[string]string{
		"red":    "#ff0001",
		"white":  "#000000",
		"yellow": "#00ff00",
		"green":  "#000fff",
		"blue":   "#00f00f",
	}
	//Declare first and assign later
	// Option -1:
	// var color map[string]string
	// Option -2:
	// color := make(map[string]string)
	//Adding key value pair in already declared Map
	// color["black"] = "#000000"
	// color["white"] = "#000000"
	//Delete a key-value from a map
	// delete(color, "white")
	// fmt.Println(color)
	printMap(color)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
