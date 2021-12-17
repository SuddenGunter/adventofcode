package input

import (
	"os"
	"strconv"
	"strings"
)

type Data struct {
	XLow, XHigh, YLow, YHigh int
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	line := strings.Split(string(file), "\n")[0]
	noPrefix := strings.TrimPrefix(line, "target area: ")
	split := strings.Split(noPrefix, ", ")
	xlow, xhigh, err := getPair(split[0])
	if err != nil {
		return Data{}, err
	}

	ylow, yhigh, err := getPair(split[1])
	if err != nil {
		return Data{}, err
	}

	return Data{
		XLow:  xlow,
		XHigh: xhigh,
		YLow:  ylow,
		YHigh: yhigh,
	}, nil
}

func getPair(str string) (int, int, error) {
	str = str[2:]
	split := strings.Split(str, "..")
	first, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, err
	}

	second, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, err
	}

	if first < second {
		return first, second, nil
	}

	return second, first, nil
}
