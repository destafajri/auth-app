package entity

type UserEntity struct{
	ID string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Phone string `json:"phone" db:"phone"`
	Role string `json:"role" db:"role"`
	Password string `json:"password" db:"password"`
}