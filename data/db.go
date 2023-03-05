package data

import (
	"database/sql"
	"favcrypto/utils"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DbClient struct {
	*sql.DB
}

var DB DbClient

func (d DbClient) InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %q", err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	fmt.Println("Successfully connected to the database!")
	DB = DbClient{db}
}

func (d DbClient) GetDataFromDB() Currencies {
	rows, err := DB.Query("SELECT * FROM Currency.data")
	if err != nil {
		log.Fatalf("Error querying database: %q", err)
	}
	defer rows.Close()

	currencies := Currencies{}
	for rows.Next() {
		var currency Currency
		if err := rows.Scan(&currency.ID, &currency.Favorite, &currency.Name, &currency.Symbol, &currency.Price); err != nil {
			log.Fatalf("Error scanning database rows: %q", err)
		}

		currencies = append(currencies, currency)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error reading database rows: %q", err)
	}

	fmt.Println("Successfully fetched data from the database!")
	return currencies
}

func (d DbClient) UpdatePrices(currencies Currencies) {
	tickersData, err := utils.FetchData()
	if err != nil {
		log.Fatalf("Error fetching data: %q", err)
	}

	newTickersData, err := utils.CleanUpData(tickersData)
	if err != nil {
		log.Fatalf("Error cleaning up data: %q", err)
	}

	newTickerPrices := make(utils.MapPrice)
	for _, data := range newTickersData {
		for key, value := range data {
			newTickerPrices[key] = value
		}
	}

	for _, currency := range currencies {
		priceStr, ok := newTickerPrices[currency.Symbol]
		if !ok {
			continue
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Fatalf("Error converting price to float: %q", err)
		}

		_, err = DB.Exec("UPDATE Currency.data SET price = $1 WHERE symbol = $2", price, currency.Symbol)
		if err != nil {
			log.Fatalf("Error updating database: %q", err)
		}
	}
	fmt.Println("Successfully updated the database!")
}

func (d DbClient) UpdateFavorite(b bool, s string) int64 {
	res, err := DB.Exec("UPDATE Currency.data SET favorite = $1 WHERE symbol = $2", b, s)
	if err != nil {
		log.Fatalf("Error updating database: %q", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error getting rows affected: %q", err)
	}

	info := "Successfully"
	if rowsAffected == 0 {
		info = "Unsuccessfully"
	}

	fmt.Printf("%s updated %d row(s) for %s value to %t\n", info, rowsAffected, s, b)
	return rowsAffected
}
