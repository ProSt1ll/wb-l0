package models

type Payment struct {
	Transaction  string `json:"transaction" db:"transaction" validate:"required"`
	RequestID    string `json:"request_id" db:"request_id" validate:"min=0"`
	Currency     string `json:"currency" db:"currency" validate:"required"`
	Provider     string `json:"provider" db:"provider" validate:"required"`
	Amount       int    `json:"amount" db:"amount" validate:"required"`
	PaymentDt    int    `json:"payment_dt" db:"payment_dt" validate:"required"`
	Bank         string `json:"bank" db:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost" validate:"required"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total" validate:"required"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee" validate:"gte=0"`
}
