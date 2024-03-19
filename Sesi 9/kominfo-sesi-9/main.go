package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = ""
)

type User struct {
	Id        uint
	Name      string
	Email     string
	CreatedAt time.Time
}

type Product struct {
	Id          uint
	Title       string
	Description sql.NullString
	CreatedAt   time.Time
}

type UserProduct struct {
	User
	Product
}

func StringOrNil(str string) interface{} {
	if str == "" {
		return nil
	}

	return str
}

func UpdateProductById(db *sql.DB, id uint, title string, description string) {
	updateProductByIdQuery := `
		UPDATE "products"
		SET title = $2,
		description = $3
		WHERE id = $1
	`

	_, err := db.Exec(updateProductByIdQuery, id, title, description)

	if err != nil {
		panic(fmt.Sprintf("update product by id: %s", err.Error()))
	}
}

func DeleteProductById(db *sql.DB, id uint) {

	deleteProductById := `
		DELETE from "products"
		WHERE id = $1
	`

	_, err := db.Exec(deleteProductById, id)

	if err != nil {
		panic(fmt.Sprintf("delete product by id: %s", err.Error()))
	}

}

func getProducts(db *sql.DB) {
	getProductsQuery := `
		SELECT id, title, description, created_at from "products"
	`

	rows, err := db.Query(getProductsQuery)

	if err != nil {
		panic(fmt.Sprintf("error getting products: %s", err.Error()))
	}

	products := []Product{}

	for rows.Next() {
		var p Product

		err = rows.Scan(&p.Id, &p.Title, &p.Description, &p.CreatedAt)

		if err != nil {
			panic(fmt.Sprintf("error while scanning product: %s\n", err.Error()))
		}

		products = append(products, p)
	}

	fmt.Printf("products => %#v\n", products)
}

func getProductById(db *sql.DB, id uint) {
	getProductByIdQuery := `
		SELECT id, title, description, created_at from "products"
		WHERE id = $1
	`

	var p Product

	err := db.QueryRow(getProductByIdQuery, id).Scan(
		&p.Id,
		&p.Title,
		&p.Description,
		&p.CreatedAt,
	)

	if err != nil {
		panic(fmt.Sprintf("error get product by id: %s", err.Error()))
	}

	fmt.Printf("product data => %+v\n", p)
}

func createProduct(db *sql.DB, title string, description string) {

	createProductQuery := `
		INSERT INTO "products"
		(title, description)
		VALUES($1, $2)
	`

	_, err := db.Exec(createProductQuery, title, StringOrNil(description))

	if err != nil {
		panic(fmt.Sprintf("error creating product: %s", err.Error()))
	}
}

func createProductWithReturning(db *sql.DB, title string, description string) {
	createProductQuery := `
		INSERT INTO "products"
		(title, description)
		VALUES($1, $2)
		RETURNING id, title, description, created_at
	`

	var p Product

	// var productId uint

	// var productTitle string

	// var productDescription sql.NullString

	// var createdAt time.Time

	err := db.QueryRow(createProductQuery, title, description).Scan(
		&p.Id, &p.Title, &p.Description, &p.CreatedAt,
	)

	if err != nil {
		panic(fmt.Sprintf("error creating product: %s", err.Error()))
	}

	// fmt.Printf("productId: (%d), title: (%s)\n", productId, productTitle)

	fmt.Printf("new product => %+v\n", p)

}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// "host=localhost port=5432 user=airelljordan password=postgres dbname=learn-pg sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	productTable := `
		CREATE TABLE IF NOT EXISTS "products"  (
			id SERIAL PRIMARY KEY,
			title varchar(255) NOT NULL,
			description varchar(255),
			created_at timestamptz NOT NULL DEFAULT now()
		)
	`

	_, err = db.Exec(productTable)

	if err != nil {
		panic(fmt.Sprintf("error creating product table: %s", err.Error()))
	}

	// createProductWithReturning(db, "Buku Atomic Habbit", "bagus")

	// getProductById(db, 6)

	// getProducts(db)

	// UpdateProductById(db, 2, "Evercoss", "Hp lama")

	// DeleteProductById(db, 5)

	// createProduct(db, "Kopi Gula Aren", "")

	// createProduct(db, "Kopi Vietnam Drip", "")
}
