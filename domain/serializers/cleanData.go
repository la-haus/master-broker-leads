package serializers

import (
	"strings"

	"github.com/la-haus/master-brokers-charge-leads/configs"
	"github.com/la-haus/master-brokers-charge-leads/domain/entity"
	"github.com/la-haus/master-brokers-charge-leads/domain/util"
	"github.com/la-haus/master-brokers-charge-leads/domain/validator"
)

func CleanLead(lead entity.Lead, config *configs.Config) entity.Lead {
	country := util.DefineCountry(lead)
	currency := util.DefineAlternativeCurrency(country)
	lead.Email = strings.ToLower(lead.Email)
	lead.BudgetResponse = util.ParseBudget(strings.ToLower(lead.Budget), currency)
	lead.ExpectedPeriod = util.ParseExpectedPurchasePeriod(strings.ToLower(lead.Expected_purchase_time))
	lead.Monthly_payment_budget = util.ParseMonthlyBudget(strings.ToLower(lead.Monthly_payment))
	lead.Preferred_property_condition = util.ParsePropertyCondition(strings.ToLower(lead.Preferred_property_condition))
	lead.Purchase_purpose = util.ParsePropertyPurpose(strings.ToLower(lead.Purchase_purpose))
	lead.Phone = validator.ValidatePhone(lead.Phone, country[:2], config)
	return lead
}
