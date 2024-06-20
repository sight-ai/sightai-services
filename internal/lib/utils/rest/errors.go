package rest

import (
	"context"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"net/http"
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// All logic layer error message should be mapped to a http error defined here:
// https://axinan.atlassian.net/wiki/spaces/TECH/pages/197951489/HTTP+Response+Codes
var AxHttpCodes = []int{
	http.StatusOK,
	http.StatusBadRequest,
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusConflict,
	http.StatusTooManyRequests,
	http.StatusInternalServerError,
}

// ErrBadRequest returns 400 error
func ErrBadRequest(msg string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		"bad_request",
		msg,
	})
}

// ErrUnauthorized returns 401 error
func ErrUnauthorized(msg string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnauthorized, struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		"unauthorized",
		msg,
	})
}

// ErrForbidden returns 403 error
func ErrForbidden(msg string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusForbidden, struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		"forbidden",
		msg,
	})
}

// ErrConflict returns 409 error
func ErrConflict(msg string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusConflict, struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		"conflict",
		msg,
	})
}

// ErrInternalServerError returns 500 error
func ErrInternalServerError(ctx context.Context, err error) *echo.HTTPError {
	log.Error(ctx).Msgf("internal server error: %+v", err)
	return echo.NewHTTPError(http.StatusInternalServerError, struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		"internal_server_error",
		err.Error(),
	})
}

// ErrInternalServerErrorWithMessage returns 500 error
func ErrInternalServerErrorWithMessage(ctx context.Context, err error, message string) *echo.HTTPError {
	log.Error(ctx).Msgf("internal server error: %+v", err)
	return echo.NewHTTPError(http.StatusInternalServerError, struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		"internal_server_error",
		message,
	})
}

// ErrFromGormError returns 500 or 400 error
func ErrFromGormError(ctx context.Context, err error, msg string) error {
	if gorm.IsRecordNotFoundError(err) {
		return ErrBadRequest(msg)
	}
	return ErrInternalServerError(ctx, err)
}

// ErrFromSwaggerClientError returns 500 or 400 error
// example err.Error()
// [POST /users/{user_id}/campaigns/{campaign_id}][400] postUsersUserIdCampaignsCampaignIdBadRequest  &{Code:bad_request Message:insufficient_points}
func ErrFromSwaggerClientError(ctx context.Context, err error) error {
	fmt.Println(err.Error())
	re := regexp.MustCompile(`Code:([^}]*) Message:([^}]*)`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) < 3 {
		return ErrInternalServerError(ctx, err)
	}
	code := matches[1]
	message := matches[2]
	if code == "internal_server_error" {
		return ErrInternalServerErrorWithMessage(ctx, err, message)
	}
	return ErrBadRequest(message)
}
