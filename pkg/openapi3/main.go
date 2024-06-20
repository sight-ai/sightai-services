package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// AdminDeposit - 
	e.POST("/v1/deposit", c.AdminDeposit)

	// UpsertGateway - 
	e.PUT("/v1/gateway", c.UpsertGateway)

	// CreateReceipt - 
	e.POST("/v1/receipt", c.CreateReceipt)

	// GatewayGetReceipts - 
	e.GET("/v1/gateways/:account_id/receipts", c.GatewayGetReceipts)

	// GatewayWithdraw - 
	e.POST("/v1/gateways/:account_id/withdraw", c.GatewayWithdraw)

	// GetAccountAllowances - 
	e.GET("/v1/accounts/:account_id/allowances", c.GetAccountAllowances)

	// GetAccountInfo - 
	e.GET("/v1/accounts/:account_id", c.GetAccountInfo)

	// GetAccountTransactions - 
	e.GET("/v1/accounts/:account_id/transactions", c.GetAccountTransactions)

	// GetGateways - 
	e.GET("/v1/gateways", c.GetGateways)

	// GetNextNonce - 
	e.GET("/v1/next_nonce", c.GetNextNonce)

	// SignAllowance - 
	e.POST("/v1/sign_allowance", c.SignAllowance)

	// SignIn - 
	e.POST("/v1/sign_in", c.SignIn)

	// UserGetReceipts - 
	e.GET("/v1/accounts/:account_id/receipts", c.UserGetReceipts)

	// Withdraw - 
	e.POST("/v1/withdraw", c.Withdraw)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}