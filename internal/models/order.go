package models

import "time"

type Order struct {
	OrderUID    string `json:"order_uid" db:"order_uid" validate:"required"`
	TrackNumber string `json:"track_number" db:"track_number" validate:"required"`
	Entry       string `json:"entry" db:"entry" validate:"required"`

	Delivery Delivery `json:"delivery" db:"delivery" validate:"required"`
	Payment  Payment  `json:"payment" db:"payment" validate:"required"`
	Items    []Item   `json:"items" db:"items" validate:"required"`

	Locale            string    `json:"locale" db:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature" validate:"min=0"`
	CustomerID        string    `json:"customer_id" db:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" db:"shardkey" validate:"required"`
	SmID              int       `json:"sm_id" db:"sm_id" validate:"required"`
	DateCreated       time.Time `json:"date_created" db:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" db:"oof_shard" validate:"required"`
}
