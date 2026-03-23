package db

import (
	"database/sql"

	"github.com/example/myapp/models"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() error {
	var err error
	DB, err = sql.Open("sqlite3", "./products.db")
	if err != nil {
		return err
	}

	return createTable()
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}

func createTable() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			price TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	var count int
	DB.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if count == 0 {
		products := []models.Product{
			{Name: "Продукт А", Description: "Высокое качество и надёжность", Price: "1000 ₽"},
			{Name: "Продукт Б", Description: "Профессиональное решение", Price: "2500 ₽"},
			{Name: "Продукт В", Description: "Премиум сегмент", Price: "5000 ₽"},
			{Name: "Продукт Г", Description: "Бюджетный вариант", Price: "750 ₽"},
			{Name: "Продукт Д", Description: "Для бизнеса", Price: "3000 ₽"},
			{Name: "Продукт Е", Description: "Универсальное решение", Price: "4500 ₽"},
		}
		for _, p := range products {
			DB.Exec("INSERT INTO products (name, description, price) VALUES (?, ?, ?)", p.Name, p.Description, p.Price)
		}
	}

	return nil
}
