package lib_subgraph

import (
	"context"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
)

func GetWithdrawEventsAfter(ctx context.Context, client *Client, lastTimestamp int64) ([]WithdrawEvent, error) {

	query := fmt.Sprintf("query{withdrawEvents(where: {timestamp_gte: %d}) {id to amount nonce timestamp}}", lastTimestamp)

	res := struct {
		WithdrawEvents []WithdrawEvent `json:"withdrawEvents"`
	}{}

	err := client.Graphql.Exec(ctx, query, &res, map[string]any{})
	if err != nil {
		log.Error(ctx).Err(err).Msgf("get subgraph error, %s client ", client.Node.ChainName)
		return nil, err
	}

	return res.WithdrawEvents, nil
}

func GetDepositEventsAfter(ctx context.Context, c *Client, lastTimestamp int64) ([]DepositEvent, error) {
	query := fmt.Sprintf("query{depositEvents(where: {timestamp_gte: %d}) {id from to amount timestamp}}", lastTimestamp)

	res := struct {
		DepositEvents []DepositEvent `json:"depositEvents"`
	}{}

	err := c.Graphql.Exec(ctx, query, &res, map[string]any{})
	if err != nil {
		log.Error(ctx).Err(err).Msgf("get subgraph error, %s client", c.Node.ChainName)
		return nil, err
	}

	return res.DepositEvents, nil
}
