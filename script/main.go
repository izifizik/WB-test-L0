package main

import (
	"WB-test-L0/internal/app"
	"WB-test-L0/internal/domain/model"
	"encoding/json"
	"log"
	"strconv"
	"time"
)

func main() {
	conn, err := app.StanConnect("pub-er", "localhost", "4222")
	if err != nil {
		log.Fatal(err.Error())
	}
	uuid := "0"
	entity := model.Entity{
		OrderUID:    uuid,
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: model.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: model.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDt:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: []model.Items{
			{
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmID:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
	}
	subj := "wb"

	for i := 0; i < 10; i++ {
		entity.OrderUID = uuid + strconv.Itoa(i)
		message, err := json.Marshal(entity)
		if err != nil {
			log.Fatalf(err.Error())
		}

		err = conn.Publish(subj, message)
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}
		log.Printf("Published [%s] : '%s'\n", subj, message)
	}
}
