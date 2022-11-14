package entity

type Farmer struct {
	ID     				string 	`json:"uuid" db:"uuid"`
	Komoditas 			string 	`json:"komoditas" db:"komoditas"`
	Area_provinsi 		string 	`json:"area_provinsi" db:"area_provinsi"`
	Area_kota			string 	`json:"area_kota" db:"area_kota"`
	Size				string	`json:"size" db:"size"`
	Price				string	`json:"price" db:"price"`
	Price_usd			string	`json:"price_usd" db:"price_usd"`
	Tgl_parsed			string	`json:"tgl_parsed" db:"tgl_parsed"`
	Timestamp			string	`json:"timestamp" db:"timestamp"`
}