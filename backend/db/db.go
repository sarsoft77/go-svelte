package db

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() error {
	var err error
	DB, err = sql.Open("sqlite3", "./data/products.db")
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
	DB.Exec("DROP TABLE IF EXISTS products")
	_, err := DB.Exec(`
		CREATE TABLE products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			price TEXT NOT NULL
		)
	`)
	DB.Exec("PRAGMA case_sensitive_like = OFF")
	if err != nil {
		return err
	}

	DB.Exec(`CREATE INDEX IF NOT EXISTS idx_products_name ON products(name)`)

	seedProducts()

	return nil
}

func seedProducts() {
	rand.Seed(time.Now().UnixNano())

	prefixes := []string{"Про", "Супер", "Мега", "Ультра", "Премиум", "Бизнес", "Эксперт", "Стандарт", "Базовый", "Элит"}
	products := []string{"Станок", "Инструмент", "Оборудование", "Система", "Устройство", "Аппарат", "Комплекс", "Решение", "Продукт", "Изделие"}
	suffixes := []string{"Экстра", "Плюс", "Про", "Люкс", "Стандарт", "Оптима", "Макси", "Мини", "Универсал", "Профи"}
	descriptions := []string{
		"Высокое качество и надёжность", "Профессиональное решение", "Премиум сегмент",
		"Бюджетный вариант", "Для бизнеса", "Универсальное решение",
		"Инновационная разработка", "Проверенное временем", "Современный дизайн",
		"Экологически чистый материал", "Энергосберегающая технология",
		"Безопасная эксплуатация", "Гарантия качества", "Сервисное обслуживание",
	}

	tx, _ := DB.Begin()
	stmt, _ := tx.Prepare("INSERT INTO products (name, description, price) VALUES (?, ?, ?)")

	batch := 1000
	for i := 0; i < 10000; i++ {
		prefix := prefixes[rand.Intn(len(prefixes))]
		product := products[rand.Intn(len(products))]
		suffix := suffixes[rand.Intn(len(suffixes))]
		desc := descriptions[rand.Intn(len(descriptions))]
		price := rand.Intn(99000) + 100

		name := fmt.Sprintf("%s %s %s", prefix, product, suffix)
		priceStr := fmt.Sprintf("%d ₽", price)

		stmt.Exec(name, desc, priceStr)

		if (i+1)%batch == 0 {
			tx.Commit()
			tx, _ = DB.Begin()
			stmt, _ = tx.Prepare("INSERT INTO products (name, description, price) VALUES (?, ?, ?)")
		}
	}

	stmt.Close()
	tx.Commit()
}
