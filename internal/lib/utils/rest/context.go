package rest

import (
	"context"

	"github.com/labstack/echo/v4"
)

func NewContext(c echo.Context) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, echo.HeaderXRequestID, c.Request().Header.Get(echo.HeaderXRequestID))
	ctx = context.WithValue(ctx, XUserToken, c.Request().Header.Get(XUserToken))
	ctx = context.WithValue(ctx, XSignature, c.Request().Header.Get(XSignature))
	return ctx
}
