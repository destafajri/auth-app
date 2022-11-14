package helper

var PRICE_IDR []string

type ResponseCurrency struct {
	Success 	bool			`json:"success"`
	Query  		interface{} 	`json:"query"`
	Info   		interface{} 	`json:"info"`
	Date		string			`json:"date"`
	Result		float64			`json:"result"`
}
