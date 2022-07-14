package util

import (
	"regexp"
	"strings"

	"github.com/la-haus/master-brokers-charge-leads/domain/entity"
)

func DefineAlternativeCurrency(country string) string {
	if strings.Contains(country, "MX") {
		return "MXN"
	} else if strings.Contains(country, "CO") {
		return "COP"
	}
	return ""
}

func DefineCurrency(budget string) string {
	if budget == "" {
		return ""
	} else {
		if strings.Contains(budget, "millones") {
			return "COP"
		}
		if strings.Contains(budget, "mil") {
			return "MXN"
		}
		if strings.Contains(budget, "mdp") {
			return "MXN"
		}
		if strings.Contains(budget, "usd") {
			return "USD"
		}
		return ""
	}
}

func DefineCountry(lead entity.Lead) string {
	if lead.LocationOfInterestCodes != "" {
		return lead.LocationOfInterestCodes
	} else if lead.Hub != "" {
		return lead.Hub
	} else {
		return ""
	}
}

func DefineUnits(budget string) (string, string) {
	regex := regexp.MustCompile(`millones|mil|mdp g`)
	splitted := regex.FindAllString(budget, -1)
	if len(splitted) == 1 {
		return "", splitted[0]
	} else if len(splitted) == 2 {
		return splitted[0], splitted[1]
	}
	return "", ""
}
