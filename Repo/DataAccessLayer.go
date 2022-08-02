package Repo

import "errors"

func GetAllTrans() Transactions {
	return InMemoryData
}

func GetSingleTrans(Id string) (error, *Transaction) {
	var transaction *Transaction
	for _, value := range InMemoryData {
		if value.Id == Id {
			transaction = &value
		}
	}
	if transaction == nil {
		return errors.New("The Transaction doesn't exist in DB"), nil
	} else {
		return nil, transaction
	}

}
func AddTransaction(trans *Transaction) (Transactions, string) {
	InMemoryData := append(Transactions{}, *trans)
	return InMemoryData, "Data was added sccessfully"
}
