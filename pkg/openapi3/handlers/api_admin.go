package handlers
import (
    "github.com/GIT_USER_ID/GIT_REPO_ID/models"
    "github.com/labstack/echo/v4"
    "net/http"
)

// AdminDeposit - 
func (c *Container) AdminDeposit(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}


// UpsertGateway - 
func (c *Container) UpsertGateway(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}
