package util

import (
	"regexp"
	"strconv"
	"strings"
)

func SplitNumbers(value string) (float64, float64) {
	regex := regexp.MustCompile(`\d+\.?\d*|\.\d+g`)
	splitted := regex.FindAllString(value, -1)
	if len(splitted) == 0 {
		return 0, 0
	} else if len(splitted) == 1 {
		if strings.Contains(value, "más") || strings.Contains(value, "mas") {
			max, err := strconv.ParseFloat(splitted[0], 64)
			if err != nil {
				return 0, 0
			}
			return 0, max
		} else if strings.Contains(value, "menos") {
			min, err := strconv.ParseFloat(splitted[0], 64)
			if err != nil {
				return 0, 0
			}
			return min, 0
		}
	} else if len(splitted) == 2 {
		min, err := strconv.ParseFloat(splitted[0], 64)
		if err != nil {
			return 0, 0
		}
		max, err := strconv.ParseFloat(splitted[1], 64)
		if err != nil {
			return 0, 0
		}
		return min, max
	}
	return 0, 0
}

func AddZerosToBudget(number float64, units string) float64 {
	if units == "" || number == 0 {
		return number
	}
	if units == "mil" {
		return number * 1000
	} else {
		return number * 1000000
	}
}
