package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}

	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			return
		}

		k, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return
		}

		minTemp := 15
		maxTemp := 30

		for j := 0; j < k; j++ {
			if !scanner.Scan() {
				return
			}
			line := scanner.Text()
			parts := strings.Fields(line)
			if len(parts) != 2 {
				continue
			}

			op := parts[0]
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}

			if op == ">=" {
				if val > minTemp {
					minTemp = val
				}
			} else if op == "<=" {
				if val < maxTemp {
					maxTemp = val
				}
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
