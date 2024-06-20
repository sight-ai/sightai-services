package handlers

import (
	"context"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/comm_utils"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/capybaralabs-xyz/sightai-services/internal/logic"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *Container) GetGatewayReceipts(ctx echo.Context) error {
	account, err := getAccountFromJWT(ctx)
	if err != nil {
		return err
	}

	if account.Role != types.AccountRoleGateway.String() {
		return rest.ErrForbidden("Only gateways have permission")
	}

	accountIDFromPath, err := getAccountIDFromPath(ctx)
	if err != nil {
		return err
	}

	if account.ID != accountIDFromPath {
		return rest.ErrBadRequest("account id not matching user token")
	}

	receiptStatus, _ := getReceiptStatusFromQuery(ctx)

	var userId *uint
	userAddress := getGatewayAddressFromQuery(ctx)
	if userAddress != "" {
		addr, err := comm_utils.ToEthAddress(userAddress)
		if err == nil {
			gateway, err := entities.AccountDao.GetByAddress(context.Background(), addr)
			if err == nil {
				userId = &gateway.ID
			}
		}
		if err != nil {
			return rest.ErrBadRequest(fmt.Sprintf("userAddress err: %s", err.Error()))
		}
	}

	page, err := getPageFromQuery(ctx)
	if err != nil {
		return err
	}

	pageSize, err := getPageSizeFromQuery(ctx)
	if err != nil {
		return err
	}

	before, _ := getBeforeFromQuery(ctx)

	after, _ := getAfterFromQuery(ctx)

	newCtx := rest.NewContext(ctx)

	rsp, err := logic.GetAccountReceipts(newCtx, &account.ID, userId, page, pageSize, receiptStatus, before, after)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (c *Container) UploadReceipt(ctx echo.Context) error {
	account, err := getAccountFromJWT(ctx)
	if err != nil {
		return err
	}

	if account.Role != types.AccountRoleGateway.String() {
		return rest.ErrForbidden("Only gateway can upload receipt")
	}

	req := &models.CreateReceiptRequest{}
	if err := rest.GenRequest(ctx, req); err != nil {
		return rest.ErrBadRequest(err.Error())
	}

	newCtx := rest.NewContext(ctx)

	rsp, err := logic.UploadReceipt(newCtx, req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}

// GatewayWithdraw - Get a signed ticket from selected receipts. GatewayWithdraw from on-chain vaults with the ticket.
func (c *Container) GatewayWithdraw(ctx echo.Context) error {
	account, err := getAccountFromJWT(ctx)
	if err != nil {
		return err
	}

	if account.Role != types.AccountRoleGateway.String() {
		return rest.ErrForbidden("Only gateways have permission")
	}

	accountIDFromPath, err := getAccountIDFromPath(ctx)
	if err != nil {
		return err
	}

	if account.ID != accountIDFromPath {
		return rest.ErrBadRequest("account id not matching user token")
	}

	sig, err := getSignature(ctx)
	if err != nil {
		return err
	}

	req := &models.GatewayWithdrawRequest{}
	if err := rest.GenRequest(ctx, req); err != nil {
		return rest.ErrBadRequest(err.Error())
	}

	newCtx := rest.NewContext(ctx)

	rsp, err := logic.GatewayWithdraw(newCtx, account, sig, req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}
