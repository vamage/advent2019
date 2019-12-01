package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var fuel  = 0
	var fuel2 = 0
	file, err := os.Open("part1.input")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass,err := strconv.Atoi(scanner.Text())

		if  err != nil {
			fmt.Fprintln(os.Stderr, "not a int ", err)
		}
		 fuel += masscalc(mass)
		 fuel2 += masscalc2(mass)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Fprintf(os.Stdout,"fuel needed is %d\n",fuel)
	fmt.Fprintf(os.Stdout,"fuel2 needed is %d\n",fuel2)
}

func masscalc(mass int)(int)  {


	return (mass/3)-2

}
func masscalc2(mass int) (int) {

	fuel := (mass / 3) - 2
	if fuel > 0 {
		return fuel + masscalc2(fuel)
	} else {
		return 0
}

}