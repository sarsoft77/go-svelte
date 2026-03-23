package handlers

import (
	"strconv"
	"strings"

	"github.com/example/myapp/db"
	"github.com/example/myapp/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	search := c.Request.URL.Query().Get("search")
	sortKey := c.DefaultQuery("sort", "id")
	sortDir := c.DefaultQuery("sortDir", "asc")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	validSortKeys := map[string]string{
		"id":          "id",
		"name":        "name",
		"description": "description",
		"price":       "CAST(price AS INTEGER)",
	}
	orderColumn, ok := validSortKeys[sortKey]
	if !ok {
		orderColumn = "id"
	}
	orderDir := "ASC"
	if strings.ToLower(sortDir) == "desc" {
		orderDir = "DESC"
	}

	whereClause := ""
	if search != "" {
		searchEscaped := strings.ReplaceAll(search, "'", "''")
		whereClause = "WHERE name LIKE '%" + searchEscaped + "%'"
	}

	var total int
	countQuery := "SELECT COUNT(*) FROM products " + whereClause
	if err := db.DB.QueryRow(countQuery).Scan(&total); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	totalPages := (total + pageSize - 1) / pageSize
	offset := (page - 1) * pageSize

	query := "SELECT id, name, description, price FROM products " + whereClause +
		" ORDER BY " + orderColumn + " " + orderDir +
		" LIMIT " + strconv.Itoa(pageSize) + " OFFSET " + strconv.Itoa(offset)

	rows, err := db.DB.Query(query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		products = append(products, p)
	}

	c.JSON(200, models.ProductResponse{
		Products:   products,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	})
}

func CreateProduct(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := db.DB.Exec("INSERT INTO products (name, description, price) VALUES (?, ?, ?)", p.Name, p.Description, p.Price)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	p.ID = int(id)
	c.JSON(201, p)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?", p.Name, p.Description, p.Price, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, p)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello from Go!"})
}
