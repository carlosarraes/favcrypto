## Fav Crypto API
##### Readme in pt-br can be found here: 
Fav Crypto is a stateless API written in native Golang (without packages) that allows users to upvote their favorite cryptocurrency. It has a containerized Postgres database, making data persistent. You can check the documentation at https://fcdocs.netlify.app/.

## Tech used
The API is hosted on Google Cloud Run and uses Golang:1.20.1-alpine3.17. The database is also run on a container and is PostgresSQL hosted on ElephantSQL.

## Database
The initial database was written in SQL. There is also a Dockerfile in /data/ so that the database can be changed from ElephantSQL to another host if desired.

```sql
CREATE SCHEMA IF NOT EXISTS Currency;

CREATE TABLE Currency.data(
    id SERIAL PRIMARY KEY,
    upvotes INT NOT NULL,
    name VARCHAR(20) NOT NULL,
    symbol VARCHAR(4) NOT NULL,
    price REAL NOT NULL DEFAULT 0
);

INSERT INTO Currency.data
  (upvotes, name, symbol, price)
VALUES
  (1, 'Klever', 'KLV', 0),
  (0, 'Bitcoin', 'BTC', 0),
  (0, 'Ethereum', 'ETH', 0),
  (0, 'Enjin', 'ENJ', 0),
  (0, 'Binance', 'BNB', 0),
  (0, 'Litecoin', 'LTC', 0),
  (0, 'Apecoin', 'APE', 0),
  (0, 'Chainlink', 'LINK', 0),
  (0, 'Dogecoin', 'DOGE', 0),
  (0, 'TRON', 'TRX', 0);
```

## Endpoints
* [GET https://gofa-4wgfen3n5q-rj.a.run.app/getcoins/](https://gofa-4wgfen3n5q-rj.a.run.app/getcoins/) 
  - Retrieves all coins in the database, along with their upvote counts and updated prices. It uses the klever.io API to fetch the prices (using the Symbol-USDT tickers).
  * Possible outcomes: 200 (Status OK), 405 (Method not allowed).
* [GET https://gofa-4wgfen3n5q-rj.a.run.app/](https://gofa-4wgfen3n5q-rj.a.run.app/) 
  - Checks if the server is running.
  * Possible outcomes: 200 (Status OK), 405 (Method not allowed).
* `PATCH https://gofa-4wgfen3n5q-rj.a.run.app/upvote/:symbol` - Increments the symbol given by one in the database. (Insomina, Postman)
  * Symbols: KLV, BTC, ETH, ENJ, BNB, LTC, APE, LINK, DOGE, TRX.
  * Possible outcomes: 200 (Status OK), 404 (Coin not found), 405 (Method not allowed).

## How to use endpoints
Use your favorite client (Insomnia or Postman) to access the endpoints. For the GET request, you can access our docs and try it there at https://fcdocs.netlify.app/.

## Tests
The 3 endpoints have been tested, and you can check them out in the handlers/test folder. I was not able to get mock to work on native Golang, so i may try a package on my next API.