package Repo

import (
	"strings"

	"github.com/google/uuid"
)

func Generate_uuid() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid

}

type Transaction struct {
	Id        string  `json:"Id"`
	Amount    float64 `json:"Amount"`
	Currency  string  `json:"Currency"`
	CreatedAt string  `json:"CreatedAt"`
}

type Transactions []Transaction

var InMemoryData = Transactions{
	Transaction{
		Id:        "6c59fbd9f3424059a70430b8e844b078",
		Amount:    50000000,
		Currency:  "CLP",
		CreatedAt: "2022-08-02T11:11:55.5465019Z",
	},
	Transaction{
		Id:        "2b7133bfcfb844179f7e19f18c77725a",
		Amount:    36457,
		Currency:  "MXN",
		CreatedAt: "2022-08-02T11:11:55.5465019Z",
	},
	Transaction{
		Id:        "5135421cf0ff47959fecd8670989807b",
		Amount:    1151.12,
		Currency:  "USD",
		CreatedAt: "2022-08-02T11:11:55.5465019Z",
	},
	Transaction{
		Id:        "c6854925270841559069812a64e288e7",
		Amount:    9965000.4,
		Currency:  "COP",
		CreatedAt: "2022-08-02T11:11:55.5465019Z",
	},
	Transaction{
		Id:        "27313a7f2b5a4281aec03b4b68090118",
		Amount:    12500,
		Currency:  "BRA",
		CreatedAt: "2022-08-02T11:11:55.5465019Z",
	},
}
