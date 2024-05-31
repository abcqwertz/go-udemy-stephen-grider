package main

import (
	"fmt"
)

func main() {
	colors2 := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#00ff00",
		"green": "#00ff00",
		"black": "#ff0000",
	}

	colors2["white"] = "#ffffff"
	delete(colors, "black")

	fmt.Println(colors2, colors)
}
