package services

import (
	"workshop/config"
	"workshop/db"
	"workshop/models"
)

var database = db.DBPurchases{}

func CreatePurchase(purchase models.Purchase) (models.Purchase, error) {
	purchase.GenerateID()
	purchase.Status = config.NEW
	return purchase, database.Save(purchase.ID, purchase)
}

func GetAllPurchases() map[string]models.Purchase {
	result := map[string]models.Purchase{}
	for k, purchase := range database.GetAll() {
		if p, ok := purchase.(models.Purchase); ok {
			result[k] = p
		}
	}
	return result
}

func GetPurchaseByID(key string) (interface{}, error) {
	if purchase, error := database.GetById(key); error != nil {
		return nil, error
	} else {
		return purchase, nil
	}
}

func UpdatePurchase(key string, purchase models.Purchase) (interface{}, error) {
	if purchase, err := database.Update(key, purchase); err != nil {
		return nil, err
	} else {
		return purchase, nil
	}

}

func DeletePurchase(key string) string {
	return database.Delete(key)
}