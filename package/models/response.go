package models

type Response struct{
	status 	int 		`json:"status"`
	message string 		`json:"message"`
	data	interface{}	`json:"data"`

}