package util

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/la-haus/master-brokers-charge-leads/domain/entity"
)

func ParseBudget(budget string, alternativeCurrency string) entity.Budget {
	if budget == "" {
		return entity.Budget{
			Min:      0,
			Max:      0,
			Currency: alternativeCurrency,
		}
	}
	min, max := SplitNumbers(budget)
	currency := DefineCurrency(budget)
	minUnits, maxUnits := DefineUnits(budget)
	if minUnits == "" {
		min = AddZerosToBudget(min, maxUnits)
		max = AddZerosToBudget(max, maxUnits)
	} else {
		min = AddZerosToBudget(min, minUnits)
		max = AddZerosToBudget(max, maxUnits)
	}

	return entity.Budget{
		Min:      min,
		Max:      max,
		Currency: currency,
	}
}

func ParseExpectedPurchasePeriod(period string) entity.ExpectedPeriod {
	if period == "" {
		return entity.ExpectedPeriod{}
	}
	timeRangesMap := map[string]string{
		"entre 3 meses y 6 meses":  "MONTH_BETWEEN_3_6",
		"entre 6 meses y 12 meses": "MONTH_BETWEEN_6_12",
		"más de 6 meses":           "MONTH_BETWEEN_6_12",
		"más de 12 meses":          "MONTH_GT_12",
		"menos de 3 meses":         "MONTH_LT_3",
		"de inmediato":             "MONTH_NOW",
	}
	periodNormalized := NormalizePeriod(period)

	min, max := SplitNumbers(periodNormalized)

	minResponse := GetExpectedTime(min)
	maxResponse := GetExpectedTime(max)

	return entity.ExpectedPeriod{
		Min:  minResponse,
		Max:  maxResponse,
		Time: timeRangesMap[periodNormalized],
	}

}

func ParseMonthlyBudget(monthlyBudget string) float64 {
	if monthlyBudget == "" {
		return 0
	}
	monthlyBudgetCleaned := strings.Replace(monthlyBudget, ".", "", -1)
	monthlyBudgetCleaned = strings.Replace(monthlyBudgetCleaned, ",", "", -1)
	regex := regexp.MustCompile(`\d+\.?\d*|\.\d+g`)
	splitted := regex.FindAllString(monthlyBudgetCleaned, -1)
	if len(splitted) == 0 || len(splitted) > 1 {
		return 0
	}
	value, err := strconv.ParseFloat(splitted[0], 64)
	if err != nil {
		return 0
	}
	return value
}

func ParsePropertyCondition(propertyCondition string) string {
	if propertyCondition == "" {
		return ""
	}
	propertyConditionsMap := map[string]string{
		"nuevo": "NEW",
		"usado": "USED",
	}
	return propertyConditionsMap[propertyCondition]

}

func ParsePropertyPurpose(propertyPurpose string) string {
	if propertyPurpose == "" {
		return ""
	}
	propertyPurposesMap := map[string]string{
		"habitar":    "LIVING",
		"invertir":   "INVESTMENT",
		"vacacionar": "VACATIONING",
	}
	return propertyPurposesMap[propertyPurpose]
}
