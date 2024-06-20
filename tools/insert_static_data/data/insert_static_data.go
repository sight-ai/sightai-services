package data

import (
	"context"
)

func InsertStaticData(ctx context.Context) error {
	insertAdmin(ctx)
	insertGateways(ctx)
	return nil
}
