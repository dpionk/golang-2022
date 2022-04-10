package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os/exec"
	"strconv"
)

func find_half(from float64, to float64) float64 {
	return math.Round((from + to) / 2)
}

func main() {

	from := 0.0
	to := 100.0
	output := ""

	log.Println("start")
	cmd := exec.Command("main_level5.exe")
	out, _ := cmd.StdoutPipe()
	in, _ := cmd.StdinPipe()
	cmd.Start()
	inwriter := bufio.NewWriter(in)
	outreader := bufio.NewReader(out)

	fmt.Println(outreader.ReadString('\n'))
	fmt.Println(outreader.ReadString('\n'))

	for output != "Brawo!\n" {
		half := find_half(from, to)
		fmt.Println(half)
		inwriter.WriteString(strconv.Itoa(int(half)) + "\n")
		inwriter.Flush()
		output, _ = outreader.ReadString('\n')
		fmt.Println(output)
		switch output {
		case "Więcej...\n":
			from = half
		case "Mniej...\n":
			to = half
		}
	}
	output, _ = outreader.ReadString('\n')
	output, _ = outreader.ReadString('\n')
	fmt.Println(output)
	inwriter.WriteString("AI\n")
	fmt.Println("AI")
	inwriter.Flush()
	output, _ = outreader.ReadString('\n')
	fmt.Println(output)
	fmt.Println("N")
	inwriter.WriteString("N\n")
	inwriter.Flush()
	output, _ = outreader.ReadString('\n')
	fmt.Println(output)
}
