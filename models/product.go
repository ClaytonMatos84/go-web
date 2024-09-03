package models

import (
	"web/db"
)

type Product struct {
	Id int
	Name string
	Description string
	Price float64
	Amount int
}

func SearchAllProducts() []Product {
	db := db.ConnectDB()
	defer db.Close()
	selectProducts, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}
	for selectProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	return products
}

func InsertProduct(name, description string, price float64, amount int)  {
	db := db.ConnectDB()
	defer db.Close()

	insertProductOnDB, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertProductOnDB.Exec(name, description, price, amount)
}

func SearchById(id int) Product {
	db := db.ConnectDB()
	defer db.Close()

	productQuery, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	for productQuery.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = productQuery.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount
	}

	return product
}

func UpdateProduct(id int, name, description string, price float64, amount int)  {
	db := db.ConnectDB()
	defer db.Close()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, amount, id)
}

func RemoveProduct(id int)  {
	db := db.ConnectDB()
	defer db.Close()

	deleteProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
}
