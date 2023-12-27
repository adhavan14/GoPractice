package main

type Product struct {
	id int
	name string
	price float64
}

func New(id int, name string, price float64) (Product) {
	return Product {
		id: id,
		name: name,
		price: price,
	}
}