CREATE SCHEMA IF NOT EXISTS Currency;

CREATE TABLE Currency.data(
    id SERIAL PRIMARY KEY,
    favorite BOOLEAN NOT NULL,
    name VARCHAR(20) NOT NULL,
    symbol VARCHAR(4) NOT NULL,
    price REAL NOT NULL DEFAULT 0
);

INSERT INTO Currency.data
  (favorite, name, symbol, price)
VALUES
  (TRUE, 'Klever', 'KLV', 0),
  (FALSE, 'Bitcoin', 'BTC', 0),
  (FALSE, 'Ethereum', 'ETH', 0),
  (FALSE, 'Enjin', 'ENJ', 0),
  (FALSE, 'Binance', 'BNB', 0),
  (FALSE, 'Litecoin', 'LTC', 0),
  (FALSE, 'Apecoin', 'APE', 0),
  (FALSE, 'Chainlink', 'LINK', 0),
  (FALSE, 'Dogecoin', 'DOGE', 0),
  (FALSE, 'TRON', 'TRX', 0);
