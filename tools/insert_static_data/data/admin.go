package data

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/comm_utils"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
)

func insertAdmin(ctx context.Context) {
	addr, err := comm_utils.ToEthAddress(config.Cfg.AdminAddress)
	checkErr(err)

	admin := &entities.Account{
		Address: addr,
		Role:    types.AccountRoleAdmin.String(),
	}
	err = entities.AccountDao.Create(ctx, admin)
	checkErr(err)

}
