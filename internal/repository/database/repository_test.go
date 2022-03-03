package database

import (
	"WB-test-L0/internal/domain/model"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"time"
)

//TestRepo_CreateEntity is the successful test for CreateEntity
func TestRepo_CreateEntity(t *testing.T) {
	//SETUP
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	entity := model.Entity{
		OrderUID:    "0",
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

	data, err := json.Marshal(entity)

	mock.ExpectBegin()
	mock.ExpectExec("insert into entity").WithArgs("0", data).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	//test
	repos := NewRepository(db)
	err = repos.CreateEntity("0", data)
	if err != nil {
		t.Errorf("error was not expected while insert data: %s", err)
	}

	//check result
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//TestRepo_CreateEntity_Failure is the failing test for CreateEntity
func TestRepo_CreateEntity_Failure(t *testing.T) {
	//SETUP
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	entity := model.Entity{
		OrderUID:    "0",
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

	data, err := json.Marshal(entity)

	mock.ExpectBegin()
	mock.ExpectExec("insert into entity").
		WithArgs("0", data).WillReturnError(fmt.Errorf("error with insert"))
	mock.ExpectRollback()

	//test
	repos := NewRepository(db)
	err = repos.CreateEntity("0", data)
	if err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	//check result
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//TestRepo_FindAllEntity is the successful test for FindAllEntities
func TestRepo_FindAllEntity(t *testing.T) {
	//SETUP
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "data"}).
		AddRow("1", "{\"id\": \"1\"}").
		AddRow("2", "{\"id\": \"2\"}")

	mock.ExpectQuery("SELECT (.+) FROM entity").WillReturnRows(rows)

	//test
	repos := NewRepository(db)
	_, err = repos.FindAllEntities()
	if err != nil {
		t.Errorf("error was not expected while select data: %s", err)
	}

	//check result
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepo_FindAllEntity_Failure(t *testing.T) {
	//SETUP
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM entity").WillReturnError(fmt.Errorf("error with select"))

	//test
	repos := NewRepository(db)
	_, err = repos.FindAllEntities()
	if err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	//check result
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
