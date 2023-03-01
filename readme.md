# Fav Cryptos

Technical requirements:

- Keep the code in Github
- The parameters received by this API at requests and responses must guarantee the typing of end user making the proper validation.
- e.g: If an input is expected as a string, it can only be received as a string.
- The HTTP status of each scenario implemented must respect the W3C Status Code standards
- e.g: 200 (Status OK), 404 (Not Found), 500 (Internal Server Error)
- The structs used with upvote model must support Marshal/Unmarshal with bson, json and struct
- The API should contain unit test of methods it uses

Extra:

- Deliver the whole solution running in some free cloud service
- Persist data into a DB of your preference
- Create an interface to interact with this backend
- Deliver a documentation on how to run and how to call your methods from your API
- Show the current price of the cryptocurrencies inside the project
