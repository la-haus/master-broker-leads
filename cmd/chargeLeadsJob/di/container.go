package di

import (
	"context"

	"github.com/la-haus/master-brokers-charge-leads/configs"
	charge_leads_job "github.com/la-haus/master-brokers-charge-leads/useCases/chargeLeadLogic"

	"go.uber.org/fx"
)

var config = fx.Options(fx.Provide(configs.NewConfig))

var uc = fx.Options(fx.Provide(charge_leads_job.NewChargeLeadsUseCase))

var invokers = fx.Invoke(
	runJob,
)

func NewApp() *fx.App {
	return fx.New(
		config,
		uc,
		invokers,
	)
}

func runJob(
	lifecycle fx.Lifecycle,
	uc charge_leads_job.ChargeLeadsUseCase,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			uc.ChargeLeads()
			return nil
		},
	})
}
