package task1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(filename string) (int, error) {
	constraints, err := parseConstraints(filename)
	if err != nil {
		return 0, err
	}

	digits := make([]int8, 14)

	for _, c := range constraints {
		if c.Diff > 0 {
			digits[c.DigitI], digits[c.DigitJ] = 9, int8(9-c.Diff)
		} else {
			digits[c.DigitI], digits[c.DigitJ] = int8(9+c.Diff), 9
		}
	}

	num := 0
	for _, d := range digits {
		num = num*10 + int(d)
	}

	return num, nil
}

func parseConstraints(filename string) ([]Constraint, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	l := 0
	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	s := stack{data: make([]Entry, 0, 7)}
	constraints := make([]Constraint, 0)

	for digit := 0; digit < 14; digit++ {
		l += 4
		line := lines[l]
		l++

		if !strings.HasPrefix(line, "div z") {
			return nil, fmt.Errorf("unexpected input on line %v: %v", l, line)
		}

		// push to stack
		if line == "div z 1" {
			l += 10
			line = lines[l]
			l++

			if !strings.HasPrefix(line, "add y") {
				return nil, fmt.Errorf("unexpected input on line %v: %v", l, line)
			}

			noPrefix := strings.TrimPrefix(line, "add y ")

			num, err := strconv.Atoi(noPrefix)
			if err != nil {
				return nil, fmt.Errorf("unexpected input on line %v: %v. internal error: %w", l, line, err)
			}

			s.push(Entry{Value: num, Digit: digit})

			l += 2

			continue
		}

		line = lines[l]
		l++

		if !strings.HasPrefix(line, "add x") {
			return nil, fmt.Errorf("unexpected input on line %v: %v", l, line)
		}

		noPrefix := strings.TrimPrefix(line, "add x ")

		num, err := strconv.Atoi(noPrefix)
		if err != nil {
			return nil, fmt.Errorf("unexpected input on line %v: %v. internal error: %w", l, line, err)
		}

		entry := s.pop()

		constraints = append(constraints, Constraint{
			DigitI: digit,
			DigitJ: entry.Digit,
			Diff:   entry.Value + num,
		})

		l += 12
	}

	return constraints, nil
}
