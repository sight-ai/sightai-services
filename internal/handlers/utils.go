package handlers

import (
	"context"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/jwt_auth"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

func getAccountIDFromJWT(c echo.Context) (uint, error) {
	uid, err := jwt_auth.GetAccountID(c.Request().Header.Get(rest.XUserToken))
	if err != nil {
		return 0, rest.ErrBadRequest(err.Error())
	}
	return uid, nil
}

func getAccountFromJWT(c echo.Context) (*entities.Account, error) {
	uid, err := jwt_auth.GetAccountID(c.Request().Header.Get(rest.XUserToken))
	if err != nil {
		return nil, rest.ErrBadRequest(err.Error())
	}

	account, err := entities.AccountDao.Get(context.Background(), uid)
	if err != nil {
		return nil, rest.ErrBadRequest(fmt.Sprintf("get user err: %s", err.Error()))
	}

	err = account.FillAllowance(context.Background())
	if err != nil {
		return nil, rest.ErrBadRequest(fmt.Sprintf("fill user allowance err: %s", err.Error()))
	}

	return account, nil
}

func getSignature(c echo.Context) (string, error) {
	sig := c.Request().Header.Get(rest.XSignature)
	if sig == "" {
		return "", rest.ErrBadRequest("missing X-SIGNATURE in header")
	}
	return sig, nil
}

func getAccountIDFromPath(c echo.Context) (uint, error) {
	str := c.Param("account_id")
	if str == "" {
		return 0, rest.ErrBadRequest("missing param account_id")
	}
	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, rest.ErrBadRequest(fmt.Sprintf("invalid account_id in path: %s", str))
	}

	return uint(u64), nil
}

func getGatewayAddressFromQuery(c echo.Context) string {
	return c.QueryParam("gateway_address")
}

func getUserAddressFromQuery(c echo.Context) string {
	return c.QueryParam("user_address")
}

func getPageFromQuery(c echo.Context) (uint, error) {
	str := c.QueryParam("page")
	if str == "" {
		return 0, rest.ErrBadRequest("missing param page")
	}
	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, rest.ErrBadRequest(fmt.Sprintf("invalid page in path: %s", str))
	}

	return uint(u64), nil
}

func getPageSizeFromQuery(c echo.Context) (uint, error) {
	str := c.QueryParam("page_size")
	if str == "" {
		return 0, rest.ErrBadRequest("missing param page_size")
	}
	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, rest.ErrBadRequest(fmt.Sprintf("invalid page_size in path: %s", str))
	}

	return uint(u64), nil
}

func getLimitFromQuery(c echo.Context) (uint, error) {
	str := c.QueryParam("limit")
	if str == "" {
		return 0, rest.ErrBadRequest("missing param limit")
	}
	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, rest.ErrBadRequest(fmt.Sprintf("invalid limit in path: %s", str))
	}

	return uint(u64), nil
}

func getAfterFromQuery(c echo.Context) (*time.Time, error) {
	str := c.QueryParam("after")
	if str == "" {
		return nil, rest.ErrBadRequest("missing param after")
	}
	after, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return nil, rest.ErrBadRequest(fmt.Sprintf("invalid after in path: %s", str))
	}

	return &after, nil
}

func getBeforeFromQuery(c echo.Context) (*time.Time, error) {
	str := c.QueryParam("before")
	if str == "" {
		return nil, rest.ErrBadRequest("missing param before")
	}
	before, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return nil, rest.ErrBadRequest(fmt.Sprintf("invalid before in path: %s", str))
	}

	return &before, nil
}

func getTransactionTypeFromQuery(c echo.Context) (*types.TransactionType, error) {
	str := c.QueryParam("type")
	return types.NewTransactionTypeFromString(str)
}

func getReceiptStatusFromQuery(c echo.Context) (*types.ReceiptStatus, error) {
	str := c.QueryParam("type")
	return types.NewReceiptStatusFromString(str)
}
