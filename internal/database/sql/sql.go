package sql

import (
	//"database/sql"
	"fmt"
	"github.com/ProSt1ll/wb-l0/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type SQL struct {
	db *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS orders
(
    order_uid          varchar PRIMARY KEY NOT NULL,
    track_number       varchar,
    entry              varchar,
    locale             varchar,
    internal_signature varchar,
    customer_id        varchar,
    delivery_service   varchar,
    shardkey           varchar,
    sm_id              bigint,
    date_created       timestamp,
    oof_shard          varchar
);

CREATE TABLE IF NOT EXISTS deliveries
(
    order_uid varchar PRIMARY KEY NOT NULL REFERENCES orders(order_uid) ON DELETE CASCADE,
    name    varchar,
    phone   varchar,
    zip     varchar,
    city    varchar,
    address varchar,
    region  varchar,
    email   varchar
);

CREATE TABLE IF NOT EXISTS payments
(
    order_uid     varchar PRIMARY KEY NOT NULL REFERENCES orders(order_uid) ON DELETE CASCADE,
    transaction   varchar,
    request_id    varchar,
    currency      varchar,
    provider      varchar,
    amount        bigint,
    payment_dt    bigint,
    bank          varchar,
    delivery_cost bigint,
    goods_total   bigint,
    custom_fee    bigint
);

CREATE TABLE IF NOT EXISTS items
(
    order_uid    varchar NOT NULL REFERENCES orders(order_uid) ON DELETE CASCADE,
    chrt_id      bigint,
    track_number varchar,
    price        bigint,
    rid          varchar,
    name         varchar,
    sale         bigint,
    size         varchar,
    total_price  bigint,
    nm_id        bigint,
    brand        varchar,
    status       bigint
);`

func New() SQL {
	db, err := sqlx.Connect("pgx", "user=ilnur dbname=ilnur sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.MustExec(schema)

	return SQL{
		db: db,
	}
}

func (sql *SQL) Save(order models.Order) error {

	_, err := sql.db.Exec("INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, "+
		"customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID,
		order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)

	if err != nil {
		return fmt.Errorf("ERROR when insert into orders: %v", err)
	}

	_, err = sql.db.Exec("INSERT INTO deliveries (order_uid, name, phone, zip, city, address, region, email) values ($1, $2, $3, $4, $5, $6, $7, $8)", order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)

	if err != nil {
		return fmt.Errorf("ERROR when insert into deliveries: %v", err)
	}

	_, err = sql.db.Exec("INSERT INTO payments (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) "+
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt,
		order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)

	if err != nil {
		return fmt.Errorf("ERROR when insert into order.Payment: %v", err)
	}

	for _, item := range order.Items {
		_, err = sql.db.Exec("INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) "+
			"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", order.OrderUID, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice,
			item.NmID, item.Brand, item.Status)
		if err != nil {
			return fmt.Errorf("ERROR when insert into items: %v", err)
		}
	}
	return nil
}

func (sql *SQL) Load(uid string) (models.Order, bool) {
	order := models.Order{}
	var items []models.Item
	payment := models.Payment{}
	delivery := models.Delivery{}
	err := sql.db.Get(&order, "SELECT * FROM orders WHERE order_uid=$1", uid)
	if err != nil {
		log.Fatal(err)
	}
	if len(order.OrderUID) == 0 {
		return order, false
	}
	err = sql.db.Select(&items, "SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items WHERE order_uid=$1", uid)
	if err != nil {
		log.Fatal(err)
	}
	err = sql.db.Get(&payment, "SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payments WHERE order_uid=$1", uid)
	if err != nil {
		log.Fatal(err)
	}
	err = sql.db.Get(&delivery, "SELECT name, phone, zip, city, address, region, email FROM deliveries WHERE order_uid=$1", uid)
	if err != nil {
		log.Fatal(err)
	}
	order.Payment = payment
	order.Items = items
	order.Delivery = delivery
	return order, true
}

func (sql *SQL) LoadAll() ([]models.Order, bool) {
	var orders []models.Order

	err := sql.db.Select(&orders, "SELECT * FROM orders")
	if err != nil {
		log.Fatal(err)
	}
	if len(orders) == 0 {
		return orders, false
	}
	for i, _ := range orders {
		orders[i], _ = sql.Load(orders[i].OrderUID)
	}
	return orders, true
}
