package categorymodel

import "crud/config"

func GetAll() {
	rows, err := config.DB.Query(`SELECT * FROM categories`) 
	
}