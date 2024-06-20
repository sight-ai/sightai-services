package services

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants"
	"github.com/capybaralabs-xyz/sightai-services/internal/handlers"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest/sight_middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net"
	"net/http"
)

func StartRest() (*echo.Echo, chan error, error) {
	errCh := make(chan error)
	e := echo.New()
	err := SetupRest(e)
	if err != nil {
		return nil, nil, err
	}
	go func() {
		errCh <- e.StartServer(e.Server)
	}()
	return e, errCh, nil
}

func StopRest(e *echo.Echo) error {
	return e.Shutdown(context.Background())
}

func SetupRest(e *echo.Echo) error {
	c, _ := handlers.NewContainer()

	e.HideBanner = true
	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	// add middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(sight_middleware.AddRequestID())
	e.Use(sight_middleware.RequestDump(config.Cfg.Env))
	e.Use(sight_middleware.ResponseDump(config.Cfg.Env))
	if config.Cfg.Env == "prod" {
		e.Use(sight_middleware.ApiLoggingJson())
		e.Use(sight_middleware.ApiMetrics(config.Cfg.AppName, config.Cfg.Env))
	} else {
		e.Use(sight_middleware.ApiLogging(config.Cfg.AppName))
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	// health check service
	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, constants.OKResponse)
	})

	// api v1, auth router and no-auth router depends on if jwt token is needed
	authRouter := e.Group("/v1")
	authRouter.Use(sight_middleware.Jwt())
	noAuthRouter := e.Group("/v1")

	// General endpoints for all accounts
	noAuthRouter.GET("/accounts/:account_id", c.GetAccount)
	noAuthRouter.GET("/gateways", c.GetGateways)
	noAuthRouter.POST("/sign_in", c.SignIn)

	authRouter.GET("/next_nonce", c.GetNextNonce)
	authRouter.GET("/accounts/:account_id/allowances", c.GetAccountAllowances)
	authRouter.GET("/accounts/:account_id/transactions", c.GetAccountTransactions)
	authRouter.GET("/accounts/:account_id/receipts", c.GetAccountReceipts)

	authRouter.POST("/withdraw", c.Withdraw)
	authRouter.POST("/sign_allowance", c.SignAllowance)

	// Gateway endpoints
	authRouter.GET("/gateways/:account_id/receipts", c.GetGatewayReceipts)
	authRouter.POST("/receipt", c.UploadReceipt)
	authRouter.POST("/gateways/:account_id/withdraw", c.GatewayWithdraw)

	// Admin endpoints
	authRouter.PUT("/gateway", c.UpsertGateway)
	authRouter.POST("/deposit", c.Deposit)

	// register server
	var err error
	e.Listener, err = net.Listen("tcp", config.Cfg.HostAndPort)
	return err
}
