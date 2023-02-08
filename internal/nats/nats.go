package nats

import (
	"encoding/json"
	"fmt"
	"github.com/ProSt1ll/wb-l0/internal/database"
	"github.com/ProSt1ll/wb-l0/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"sync"
)

func Subscribe(db database.Database) {
	validate := validator.New()
	sub, err := stan.Connect("wb", "sub")

	if err != nil {
		fmt.Printf("ERROR: subscriber can't connect to nats: %v\n", err)
		return
	}

	defer func() {
		err := sub.Close()
		if err != nil {

		}
	}()

	func() {

		_, err := sub.Subscribe("order", func(m *stan.Msg) {
			var order models.Order

			err := json.Unmarshal(m.Data, &order)
			if err != nil {
				fmt.Printf("ERROR: can't unmarshal json: %v\n", err)
				return
			}

			err = validate.Struct(order)
			if err != nil {
				fmt.Printf("ERROR: can't validate order: %v\n", err)
				return
			}

			err = db.Save(order)

			if err != nil {
				fmt.Printf("ERROR: can't saver order: %v\n", err)
				return
			}

		})
		if err != nil {
			fmt.Printf("ERROR: can't subscribe : %v\n", err)
			return
		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
