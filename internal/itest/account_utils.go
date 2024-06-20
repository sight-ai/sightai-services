package itest

import (
	"context"
	openapi "github.com/capybaralabs-xyz/sightai-services/internal/itest/client"
)

func Deposit(jwt, address, amount string) (*openapi.SimpleMessageResponse, error) {
	request := openapi.DepositRequest{
		Address: address,
		Amount:  amount,
	}

	ctx := GetContextWithUserToken(context.Background(), jwt)

	deposit := TestContext.client.UserApi.AdminDeposit(ctx)
	deposit = deposit.DepositRequest(request)
	resp, _, err := deposit.Execute()

	if err != nil {
		return nil, err
	} else if resp != nil {
		return resp, nil
	}

	return nil, nil
}
