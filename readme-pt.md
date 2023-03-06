# Fav Crypto API
Fav Crypto é uma stateless API feita em Golang nativo (sem pkgs) que permite que os usuários upvote suas criptomoedas favoritas. Ele tem um banco de dados PostgresSQL em um contêiner, então os dados são persistentes! Você pode verificar a documentação neste link: https://fcdocs.netlify.app/

## Tecnologias utilizadas
A API é hospedada no Google Cloud Run, utilizando Golang:1.20.1-alpine3.17. O banco de dados também é executado em um contêiner, hospedado no ElephantSQL.

## Banco de Dados
O banco de dados inicial foi escrito em .sql, há também um arquivo Dockerfile em /data/, assim tenho a opção de alterar o banco de dados do ElephantSQL para outro host.

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

GET https://gofa-4wgfen3n5q-rj.a.run.app/getcoins/ - Recupera todas as moedas do banco de dados, com suas contagens de votos e preços atualizados, usa a API klever.io para buscar os preços. (Usa os tickers Symbol-USDT para filtrar os dados.)
*Possíveis resultados: 200 (StatusOk), 405 (Method not allowed)

GET https://gofa-4wgfen3n5q-rj.a.run.app/ - Verifica se o servidor está em execução.
*Possíveis resultados: 200 (StatusOk), 405 (Method not allowed)

PATCH https://gofa-4wgfen3n5q-rj.a.run.app/upvote/:symbol - Incrementa no banco de dados o símbolo fornecido em um.
Símbolos: KLV, BTC, ETH, ENJ, BNB, LTC, APE, LINK, DOGE, TRX
*Possíveis resultados: 200 (StatusOk), 404 (Coin not found), 405 (Method not allowed)

## Como usar os endpoints
Use seu app favorito (Insomnia ou Postman) para acessar os endpoints! Para a solicitação GET, você pode acessar a documentação e testar por lá (https://fcdocs.netlify.app/)

## Testes
Os 3 endpoints foram testados e você pode verificar no diretório handlers/test! Não consegui fazer o mock funcionar no Golang nativo, talvez tente um pacote na minha próxima API.