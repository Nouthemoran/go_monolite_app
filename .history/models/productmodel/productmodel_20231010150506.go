package productmodel

import "crud/entities"

func GetAll() []entities.Product {
	config.DB.Query(`
		SELECT
			prod
	`)
}