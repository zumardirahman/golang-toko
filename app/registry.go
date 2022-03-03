package app

//guna registry ini untuk mengumpulkan model2 ke sebuah variabel

import "github.com/zumardirahman/golang-toko/app/models"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model { //method untuk meregisterkan sebuah model
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Product{}},
		{Model: models.Category{}},
		{Model: models.Section{}},
		{Model: models.ProductImage{}},
		{Model: models.Order{}},
		{Model: models.OrderItem{}},
		{Model: models.OrderCustomer{}},
		{Model: models.Payment{}},
		{Model: models.Shipment{}},
	}

} //dipanggil dari server.go di methode initializedb
