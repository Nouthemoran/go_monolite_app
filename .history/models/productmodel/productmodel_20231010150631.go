package productmodel

import "crud/entities"

func GetAll() []entities.Product {
	config.DB.Query(`
		SELECT
			products.id,
			products.name,
			categories.name as category_name
			products.stock,
			products.description,
			products.created_at,
			products.updated_at
	`)
}