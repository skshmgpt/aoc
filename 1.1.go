package main

import (
	"bufio"
	"os"
	"strconv"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ceilDiv(a, b int) int {
	if a >= 0 {
		return (a + b - 1) / b
	}
	return a / b
}

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	var prevPos = 50
	var ans = 0

	for scanner.Scan() {
		move := scanner.Text()
		movbytes := []byte(move)
		dir := string(movbytes[:1])
		turns, _ := strconv.Atoi(string(movbytes[1:]))

		if dir == "L" {
			turns = -turns
		}

		newPos := prevPos + turns

		ans += abs(newPos) / 100
		if newPos*prevPos <= 0 {
			ans++
		}
		if prevPos == 0 {
			ans--
		}
		prevPos = newPos % 100
	}
	println(ans)
}
