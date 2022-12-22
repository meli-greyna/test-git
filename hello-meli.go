package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var pressure, _ = strconv.ParseFloat(os.Args[1], 64)
	var humidity, _ = strconv.ParseFloat(os.Args[2], 64)
	var temperature, _ = strconv.ParseFloat(os.Args[3], 64)

	fmt.Println(pressure, humidity, temperature)
}
