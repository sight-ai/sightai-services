package itest

import (
	"context"
	openapi "github.com/capybaralabs-xyz/sightai-services/internal/itest/client"
)

func SignIn(walletAddress string) (*openapi.SignInResponse, error) {
	request := openapi.SignInRequest{
		WalletAddress: walletAddress,
	}

	signIn := TestContext.client.UserApi.SignIn(context.Background())
	signIn = signIn.SignInRequest(request)
	resp, _, err := signIn.Execute()

	if err != nil {
		return nil, err
	} else if resp != nil {
		return resp, nil
	}

	return nil, nil
}
