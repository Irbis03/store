package internal

type User struct {
	Id int `json:"id"`
   	Email string `json:"email"`
   	Password string `json:"password"`
   	Role string `json:"role"`
}

// type Basket struct {
// 	id int
// 	user_id int
// }

// type BasketDevice struct {
// 	id int
// 	device_id int
// 	basket_id int
// }

type Device struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Img string `json:"img"`
}