package data

import (
	"context"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/comm_utils"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/gocarina/gocsv"
	"github.com/jinzhu/gorm"
	"os"
)

type Gateway struct {
	ID       uint   `csv:"GatewayID"`
	Address  string `csv:"Address"`
	Endpoint string `csv:"Endpoint"`
	Name     string `csv:"Name"`
}

var gateways []*Gateway

func parseProductsCsv(file string) {
	gatewaysFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	if err := gocsv.UnmarshalFile(gatewaysFile, &gateways); err != nil { // Load clients from file
		panic(err)
	}
}

func insertGateways(ctx context.Context) {
	parseProductsCsv(
		fmt.Sprintf("%s/tools/insert_static_data/data/csv/%s/gateway.csv",
			comm_utils.GetGoBaseDirectory(), config.Cfg.Env))

	for _, p := range gateways {
		addr, err := comm_utils.ToEthAddress(p.Address)
		checkErr(err)

		gateway := &entities.Gateway{
			Model: gorm.Model{
				ID: p.ID,
			},
			Address:  addr,
			Endpoint: p.Endpoint,
			Name:     p.Name,
		}

		err = entities.GatewayDao.Upsert(ctx, gateway)
		checkErr(err)

		account := &entities.Account{
			Address: addr,
			Role:    types.AccountRoleGateway.String(),
		}
		err = entities.AccountDao.Create(ctx, account)
		checkErr(err)
	}
}
