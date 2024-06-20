package handlers

import (
	"context"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/comm_utils"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/capybaralabs-xyz/sightai-services/internal/logic"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *Container) SignIn(ctx echo.Context) error {
	sig, err := getSignature(ctx)
	if err != nil {
		return err
	}

	req := &models.SignInRequest{}
	if err := rest.GenRequest(ctx, req); err != nil {
		return rest.ErrBadRequest(err.Error())
	}

	newCtx := rest.NewContext(ctx)

	rsp, err := logic.SignIn(newCtx, req, sig)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (c *Container) GetAccount(ctx echo.Context) error {
	accountID, err := getAccountIDFromPath(ctx)
	if err != nil {
		return err
	}

	newCtx := rest.NewContext(ctx)

	rsp, err := logic.GetAccount(newCtx, accountID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (c *Container) GetGateways(ctx echo.Context) error {
	newCtx := rest.NewContext(ctx)

	rsp, err := logic.GetGateways(newCtx)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (c *Container) GetNextNonce(ctx echo.Context) error {
	account, err := getAccountFromJWT(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.AccountNextNonceResponse{NextNonce: int32(account.Nonce + 1)})
}

func (c *Container) GetAccountTransactions(ctx echo.Context) error {
	accountIDFromJwt, err := getAccountIDFromJWT(ctx)
	if err != nil {
		return err
	}

	accountIDFromPath, err := getAccountIDFromPath(ctx)
	if err != nil {
		return err
	}

	if accountIDFromJwt != accountIDFromPath {
		return rest.ErrBadRequest("account id not matching user token")
	}

	transactionType, _ := getTransactionTypeFromQuery(ctx)

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

	rsp, err := logic.GetAccountTransactions(newCtx, accountIDFromJwt, page, pageSize, transactionType, before, after)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (c *Container) GetAccountAllowances(ctx echo.Context) error {
	accountIDJwt, err := getAccountIDFromJWT(ctx)
	if err != nil {
		return err
	}

	accountID, err := getAccountIDFromPath(ctx)
	if err != nil {
		return err
	}

	if accountIDJwt != accountID {
		return rest.ErrBadRequest("account id not matching user token")
	}

	newCtx := rest.NewContext(ctx)

	rsp, err := logic.GetAccountAllowances(newCtx, accountID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (c *Container) GetAccountReceipts(ctx echo.Context) error {
	accountIDFromJwt, err := getAccountIDFromJWT(ctx)
	if err != nil {
		return err
	}

	accountIDFromPath, err := getAccountIDFromPath(ctx)
	if err != nil {
		return err
	}

	if accountIDFromJwt != accountIDFromPath {
		return rest.ErrBadRequest("account id not matching user token")
	}

	receiptStatus, _ := getReceiptStatusFromQuery(ctx)

	var gatewayId *uint
	gatewayAddress := getGatewayAddressFromQuery(ctx)
	if gatewayAddress != "" {
		addr, err := comm_utils.ToEthAddress(gatewayAddress)
		if err == nil {
			gateway, err := entities.AccountDao.GetByAddress(context.Background(), addr)
			if err == nil {
				gatewayId = &gateway.ID
			}
		}
		if err != nil {
			return rest.ErrBadRequest(fmt.Sprintf("gatewayAddress err: %s", err.Error()))
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

	rsp, err := logic.GetAccountReceipts(newCtx, gatewayId, &accountIDFromJwt, page, pageSize, receiptStatus, before, after)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}

// Withdraw - Get a signed ticket. Withdraw from on-chain vaults with the ticket.
func (c *Container) Withdraw(ctx echo.Context) error {
	account, err := getAccountFromJWT(ctx)
	if err != nil {
		return err
	}

	sig, err := getSignature(ctx)
	if err != nil {
		return err
	}

	req := &models.WithdrawRequest{}
	if err := rest.GenRequest(ctx, req); err != nil {
		return rest.ErrBadRequest(err.Error())
	}

	newCtx := rest.NewContext(ctx)

	rsp, err := logic.Withdraw(newCtx, account, sig, req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (c *Container) SignAllowance(ctx echo.Context) error {
	req := &models.SignAllowanceRequest{}
	if err := rest.GenRequest(ctx, req); err != nil {
		return rest.ErrBadRequest(err.Error())
	}

	account, err := getAccountFromJWT(ctx)
	if err != nil {
		return err
	}

	newCtx := rest.NewContext(ctx)

	rsp, err := logic.SignAllowance(newCtx, account, req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}
