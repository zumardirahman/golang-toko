package models

type Model struct {
	Model interface{}
}

func RegisterModel() []Model { //method untuk meregisterkan sebuah model
	return []Model{
		{Model: User{}},
		{Model: Address{}},
		{Model: Product{}},
		{Model: Category{}},
		{Model: Section{}},
		{Model: ProductImage{}},
		{Model: Order{}},
		{Model: OrderItem{}},
		{Model: OrderCustomer{}},
		{Model: Payment{}},
		{Model: Shipment{}},
	}

} //dipanggil dari server.go di methode initializedb
