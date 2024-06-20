package models

type GatewayWithdrawRequest struct {

	ReceiptIds []int64 `json:"receipt_ids"`
}
