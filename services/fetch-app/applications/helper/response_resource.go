package helper

type ResponseFetch struct {
	ID     				string 	`json:"uuid"`
	Komoditas 			string 	`json:"komoditas"`
	Area_provinsi 		string 	`json:"area_provinsi"`
	Area_kota			string 	`json:"area_kota"`
	Size				string	`json:"size"`
	Price				string	`json:"price"`
	Tgl_parsed			string	`json:"tgl_parsed"`
	Timestamp			string	`json:"timestamp"`
}