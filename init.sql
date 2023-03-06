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
