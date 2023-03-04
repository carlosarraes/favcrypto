package main

import (
	"favcrypto/data"
)

func init() {
	currencies := data.GetDataFromDB()
	data.UpdateDataInDB(currencies)
}

func main() {
}
