package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	var ans = 0

	scanner.Scan()
	idRanges := strings.Split(scanner.Text(), ",")
	for _, idRange := range idRanges {
		boundaries := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(boundaries[0])
		end, _ := strconv.Atoi(boundaries[1])
		if len([]byte(boundaries[0]))%2 == 1 && len([]byte(boundaries[1]))%2 == 1 {
			continue
		}
		// loop can be optimized here to avoid unnecessary checks
		for i := start; i <= end; i++ {
			num := []byte(strconv.Itoa(i))
			numLen := len(num)
			if numLen%2 == 1 {
				continue
			}
			f := num[:numLen/2]
			s := num[numLen/2:]
			fnum, _ := strconv.Atoi(string(f))
			snum, _ := strconv.Atoi(string(s))
			if fnum == snum {
				ans += i
			}
		}
	}
	println(ans)
}
