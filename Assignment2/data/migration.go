package data

import (
	"assignment2/model"
	"fmt"
)

func Migrate() {
	fmt.Println("Proses Migrate Database...")

	DB.AutoMigrate(&model.Order{}, &model.Item{})

	fmt.Println("Sukses Migrate Database...")
}
