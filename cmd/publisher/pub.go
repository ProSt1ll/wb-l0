package main

import (
	"encoding/json"
	"fmt"
	"github.com/ProSt1ll/wb-l0/internal/models"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"math/rand"
	"strconv"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	rand.Seed(time.Now().UnixNano())
	sc, err := stan.Connect("wb", "pub")

	if err != nil {
		fmt.Printf("ERROR: publisher can't connect to nats: %v\n", err)
		return
	}

	defer sc.Close()
	for true {
		new_order := models.Order{
			OrderUID:    uuid.New().String(),
			TrackNumber: RandIntString(6, 9),
			Entry:       RandString(6),
			Delivery: models.Delivery{
				Name:    RandString(6),
				Phone:   "+7" + RandIntString(10, 9),
				Zip:     RandString(6),
				City:    RandString(6),
				Address: RandString(6),
				Region:  RandString(6),
				Email:   RandString(6) + "@gmail.com",
			},
			Payment: models.Payment{
				Transaction:  RandString(6),
				RequestID:    RandString(6),
				Currency:     RandString(6),
				Provider:     RandString(6),
				Amount:       rand.Intn(10000),
				PaymentDt:    rand.Intn(10000),
				Bank:         RandString(6),
				DeliveryCost: rand.Intn(10000),
				GoodsTotal:   rand.Intn(10000),
				CustomFee:    rand.Intn(10000),
			},
			Items: []models.Item{
				{
					ChrtID:      rand.Intn(10000),
					TrackNumber: RandString(6),
					Price:       rand.Intn(10000),
					Rid:         RandString(6),
					Name:        RandString(6),
					Sale:        rand.Intn(10000),
					Size:        RandString(6),
					TotalPrice:  rand.Intn(10000),
					NmID:        rand.Intn(10000),
					Brand:       RandString(6),
					Status:      rand.Intn(10000),
				},
			},
			Locale:            RandString(6),
			InternalSignature: RandString(6),
			CustomerID:        RandString(6),
			DeliveryService:   RandString(6),
			Shardkey:          RandString(6),
			SmID:              rand.Intn(10000),
			DateCreated:       time.Now(),
			OofShard:          RandString(6),
		}

		json_order, err := json.Marshal(new_order)
		if err != nil {
			fmt.Printf("ERROR: can't marshal order to json: %v\n", err)
			return
		}
		err = sc.Publish("order", json_order)
		if err != nil {
			fmt.Printf("ERROR: can't publish message: %v\n", err)
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandIntString(size int, max int) string {
	var s string
	for i := 0; i < size; i++ {
		s = s + strconv.Itoa(rand.Intn(max))
	}
	return s
}
