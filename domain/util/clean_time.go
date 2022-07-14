package util

import (
	"regexp"
	"time"
)

func GetExpectedTime(countMonths float64) string {
	if countMonths == 0 {
		return ""
	}
	date := time.Now()
	datePlusMonths := date.AddDate(0, int(countMonths), 0)
	return datePlusMonths.Format("2006-01-02")
}

func NormalizePeriod(period string) string {
	regularExp := regexp.MustCompile(`1 ano|1 año|un año|un ano`)
	return regularExp.ReplaceAllString(period, "12 meses")
}
