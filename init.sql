CREATE SCHEMA IF NOT EXISTS Currency;

CREATE TABLE Currency.data(
    id SERIAL PRIMARY KEY,
    favorite BOOLEAN NOT NULL,
    name VARCHAR(20) NOT NULL,
    symbol VARCHAR(4) NOT NULL,
    price DECIMAL(10,2)
);

INSERT INTO Currency.data
  (favorite, name, symbol, price)
VALUES
  (TRUE, 'Klever', 'KLV', NULL),
  (FALSE, 'Bitcoin', 'BTC', NULL),
  (FALSE, 'Ethereum', 'ETH', NULL),
  (FALSE, 'Enjin', 'ENJ', NULL),
  (FALSE, 'Binance', 'BNB', NULL),
  (FALSE, 'Litecoin', 'LTC', NULL),
  (FALSE, 'Apecoin', 'APE', NULL),
  (FALSE, 'Chainlink', 'LINK', NULL),
  (FALSE, 'Dogecoin', 'DOGE', NULL),
  (FALSE, 'TRON', 'TRX', NULL);
