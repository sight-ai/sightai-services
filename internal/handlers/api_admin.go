package handlers

import (
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/capybaralabs-xyz/sightai-services/internal/logic"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *Container) Deposit(ctx echo.Context) error {
	account, err := getAccountFromJWT(ctx)
	if err != nil {
		return err
	}

	if account.Role != types.AccountRoleAdmin.String() {
		return rest.ErrForbidden("Only admin can deposit")
	}

	req := &models.DepositRequest{}
	if err := rest.GenRequest(ctx, req); err != nil {
		return rest.ErrBadRequest(err.Error())
	}

	newCtx := rest.NewContext(ctx)

	rsp, err := logic.Deposit(newCtx, req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (c *Container) UpsertGateway(ctx echo.Context) error {
	account, err := getAccountFromJWT(ctx)
	if err != nil {
		return err
	}

	if account.Role != types.AccountRoleAdmin.String() {
		return rest.ErrForbidden("Only admin can deposit")
	}

	req := &models.UpsertGatewayRequest{}
	if err := rest.GenRequest(ctx, req); err != nil {
		return rest.ErrBadRequest(err.Error())
	}

	newCtx := rest.NewContext(ctx)

	rsp, err := logic.UpdateGateway(newCtx, req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rsp)
}
