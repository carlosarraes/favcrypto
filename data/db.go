package data

import (
	"database/sql"
	"favcrypto/services"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Currency struct {
	ID       int
	Favorite bool
	Name     string
	Symbol   string
	Price    float64
}

type Currencies []Currency

func ConnectToDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %q", err)
	}

	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %q", err)
	}

	fmt.Println("Successfully connected to the database!")

	return db
}

func GetDataFromDB() Currencies {
	db := ConnectToDB()

	rows, err := db.Query("SELECT * FROM Currency.data")
	if err != nil {
		log.Fatalf("Error querying database: %q", err)
	}
	defer rows.Close()

	currencies := Currencies{}
	for rows.Next() {
		var currency Currency
		err := rows.Scan(&currency.ID, &currency.Favorite, &currency.Name, &currency.Symbol, &currency.Price)
		if err != nil {
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

func UpdateDataInDB(currencies Currencies) {
	tickersData, err := services.FetchData()
	if err != nil {
		log.Fatalf("Error fetching data: %q", err)
	}

	newTickersData, err := services.CleanUpData(tickersData)
	if err != nil {
		log.Fatalf("Error cleaning up data: %q", err)
	}

	newTickerPrices := make(services.MapPrice)
	for _, data := range newTickersData {
		for key, value := range data {
			newTickerPrices[key] = value
		}
	}

	db := ConnectToDB()
	for _, currency := range currencies {
		priceStr, ok := newTickerPrices[currency.Symbol]
		if !ok {
			continue
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Fatalf("Error converting price to float: %q", err)
		}

		_, err = db.Exec("UPDATE Currency.data SET price = $1 WHERE symbol = $2", price, currency.Symbol)
		if err != nil {
			log.Fatalf("Error updating database: %q", err)
		}
	}
	fmt.Println("Successfully updated the database!")
}
